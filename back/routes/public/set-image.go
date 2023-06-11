package public

import (
	"strconv"
	"net/http"

	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
)

func setImage(path string, num string, user string, r *http.Request) {
	n, err := strconv.ParseInt(num, 10, 32)
	log.Logger.Errorf("%+v", err)

	payload := payloads.Image{
		Owner: user,
		Path: path, 
		Number: int(n),
	}

	if err := payload.Push(r); err != nil {
		log.Logger.Errorf("%+v", err)
		return
	}
}
