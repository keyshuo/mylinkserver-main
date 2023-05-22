package user

const (
	Hadlogined = "true"
	Nologin    = "false"
)

type User struct {
	Account  string `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTToken struct {
	Token string `json:"token"`
}

var jwtKey = []byte("my_secret_key") //后期改为本地字符串
