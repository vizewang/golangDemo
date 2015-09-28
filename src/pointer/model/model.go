package model
import "strconv"
type User struct {
	Username string
	password string
	age int32
}
func (this *User)String()string  {
	return "username: "+this.Username+" password: "+this.password+" age:"+strconv.Itoa(int(this.age))
}
