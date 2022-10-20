package server

func Init() {
	r := NewRouter()
  r.Run("127.0.0.1:8000")
}

