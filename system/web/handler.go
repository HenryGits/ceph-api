package web

type ConnConfig struct {
	Monitors string `json:"id" example:"192.168.113.215:6789,192.168.113.216:6789,192.168.113.217:6789" format:"Ceph Mgr ipaddr"`
	User     string `json:"user" example:"admin" format:"用户名"`
	Key      string `json:"key" example:"AQDzrN1eAUaLOhAAQ9y9mrYUPH1DqBYvrvtGug==" format:"Mgr 密钥Key"`
}
