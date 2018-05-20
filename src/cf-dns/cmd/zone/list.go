package zone

type Zone struct {
	Id   string
	Name string
}

func List() {

	// curl -X GET "https://api.cloudflare.com/client/v4/zones?name=example.com&status=active&page=1&per_page=20&order=status&direction=desc&match=all" \
	//  -H "X-Auth-Email: user@example.com" \
	//  -H "X-Auth-Key: c2547eb745079dac9320b638f5e225cf483cc5cfdda41" \
	//  -H "Content-Type: application/json"

}
