#!/bin/bash
REGION=ap-southeast-1
AWS_ACCESS_KEY_ID=dummy
AWS_ACCESS_ACCESS_KEY=dummy
OUTPUT=table
PROFILE=test-profile
HOST=http://localhost:4566
SNS_TOPICS=(
  "brick-transfer"
)
SQS_QUEUES=(
  "transfer-deduct-balance-queue"
  "proceed-transfer-queue"
  "reject-transfer-queue"
  "transfer-reversal-balance-queue"
)
SNS_SUBSCRIPTIONS=(
  "brick-transfer,accepted,transfer-deduct-balance-queue"
  "brick-transfer,proceed,proceed-transfer-queue"
  "brick-transfer,reject,reject-transfer-queue"
  "brick-transfer,reversal,transfer-reversal-balance-queue"
)

# init
aws configure set aws_access_key_id "$AWS_ACCESS_KEY_ID" --profile $PROFILE
aws configure set aws_secret_access_key "$AWS_ACCESS_ACCESS_KEY" --profile $PROFILE
aws configure set region "$REGION" --profile $PROFILE
aws configure set output "$TABLE" --profile $PROFILE

# setup topic
for SNS_TOPIC in "${SNS_TOPICS[@]}"; do
  aws --endpoint-url=$HOST sns create-topic \
      --name $SNS_TOPIC \
      --region $REGION \
      --profile $PROFILE \
      --output $OUTPUT | cat
done

# setup queue
for SQS_QUEUE in "${SQS_QUEUES[@]}"; do
  aws --endpoint-url=$HOST sqs create-queue \
      --queue-name $SQS_QUEUE \
      --profile $PROFILE \
      --region $REGION \
      --output $OUTPUT | cat
done

# setup subscription
for SNS_SUBSCRIPTION in "${SNS_SUBSCRIPTIONS[@]}"
do
  # Split the tokens
  IFS=',' read -ra T_S <<< "$SNS_SUBSCRIPTION"
  TOPIC=${T_S[0]}
  ROUTING_KEY=${T_S[1]}
  QUEUE=${T_S[2]}
  aws --endpoint-url=$HOST sns subscribe \
      --topic-arn arn:aws:sns:$REGION:000000000000:$TOPIC \
      --profile $PROFILE  \
      --protocol sqs \
      --notification-endpoint "arn:aws:sns:$REGION:000000000000:$QUEUE" \
      --attributes '{"FilterPolicy":"{\"action\":[\"'"$ROUTING_KEY"'\"]}", "RawMessageDelivery":"True"}'
      --output $OUTPUT | cat
done

echo "Localstack init script executed successfully"

# Cheatsheet
## Command to make SQS subscribe to SNS
aws --endpoint-url="http://localhost:4566" sns subscribe \
    --topic-arn arn:aws:sns:ap-southeast-1:000000000000:brick-transfer \
    --profile test-profile  \
    --protocol sqs \
    --notification-endpoint "arn:aws:sns:ap-southeast-1:000000000000:transfer-deduct-balance-queue" \
    --attributes '{"FilterPolicy":"{\"action\":[\"accepted\"]}"}' \
    --output table | cat

## Publish message
aws --endpoint-url=http://localhost:4566 sns publish \
    --region ap-southeast-1 --profile test-profile --topic-arn arn:aws:sns:ap-southeast-1:000000000000:brick-transfer \
    --message="Your message content c" \
    --message-attributes '{"action":{"DataType":"String","StringValue":"accepted"}}'

## Check message in Queue
aws --endpoint-url=http://localhost:4566 sqs receive-message \
    --queue-url http://localhost:4566/000000000000/transfer-deduct-balance-queue \
    --profile test-profile \
    --region ap-southeast-1 \
    --attribute-names All --message-attribute-names All \
    --output json | cat
### additional --query='Messages[?MessageAttributes.action.StringValue == `accepted`]' --max-number-of-messages 1 \

## Purge message in Queue
aws --endpoint-url=http://localhost:4566 sqs purge-queue \
    --queue-url http://localhost:4566/000000000000/transfer-deduct-balance-queue \
    --profile test-profile \
    --region ap-southeast-1 \
    --output json | cat

# Check all message in Queue via curl
curl "http://localhost:4566/_aws/sqs/messages?QueueUrl=http://localhost:4566/000000000000/transfer-deduct-balance-queue"

