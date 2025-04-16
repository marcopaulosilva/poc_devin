set -e

echo "Initializing LocalStack RDS..."

awslocal rds create-db-instance \
    --db-instance-identifier champions-db \
    --db-instance-class db.t3.micro \
    --engine postgres \
    --master-username postgres \
    --master-user-password postgres \
    --allocated-storage 20 \
    --db-name champions

echo "RDS instance created successfully!"

echo "Waiting for RDS instance to be available..."
awslocal rds wait db-instance-available --db-instance-identifier champions-db

echo "LocalStack RDS initialization completed!"
