package canonical

type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}