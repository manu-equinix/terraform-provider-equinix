cd fabric-gateway-test
echo "=============Initialize FG Resource==========="
terraform init
echo "=============Create FG ==========="
terraform apply -auto-approve -var="fg_name=terra_fg-demo"
echo "=============GET FG ==========="
terraform show
echo "=============Update FG Name==========="
terraform apply -refresh -auto-approve -var="fg_name=terra_fg-demo-update"
terraform show
result=$(terraform show | grep fg_result)
echo $result
id=$(echo $result | awk -F "=" '{print $2}')
echo "=============Create FG2port Connection ==========="
cd ../fg2port-test
terraform init
terraform apply -auto-approve -var="fg_uuid=$id"
echo "=============GET Connection ==========="
terraform show

echo "DONE"

