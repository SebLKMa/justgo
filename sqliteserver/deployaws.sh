# for security reason, pem file is not checked in
scp -i ~/dev/aws/martin/free-instance-seb.pem survey ubuntu@ec2-46-137-199-71.ap-southeast-1.compute.amazonaws.com:~/staging/sqliteserver/survey
scp -i ~/dev/aws/martin/free-instance-seb.pem -r templates ubuntu@ec2-46-137-199-71.ap-southeast-1.compute.amazonaws.com:~/staging/sqliteserver
