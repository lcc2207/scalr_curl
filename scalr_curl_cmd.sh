export SCALR_SERVER_URL=http://demo.scalr.club
export SCALR_API_KEY_ID=xxxx
export METHOD=GET
export QUERY_PATH='/api/v1beta0/user/6/farms/'
export SCALR_SECRET=xxxx

info=$(./scalr_curl)

echo $info | jq
