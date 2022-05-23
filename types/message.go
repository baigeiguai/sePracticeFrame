package types
type RegisterMessage int16
const (
	RegisterSuccess RegisterMessage =0
	RegisterNameExisted RegisterMessage =1
	RegisterError RegisterMessage =2
)
func (p RegisterMessage) String()string {
	switch (p){
	case RegisterSuccess: return "Register Successfully!"
	case RegisterNameExisted: return "This username has existed!"
	case RegisterError: return "Fail to Register the user!"
	default: return "Unknow"
	}
}
type LoginMessage int16 
const (
	LoginNameNotExist LoginMessage = 0 
	LoginPasswordNotCorrect LoginMessage = 1
	LoginSuccess LoginMessage =2
	LoginError LoginMessage =3
)
func (p LoginMessage)String()string{
	switch (p){
	case LoginNameNotExist : return "The username doesn't exist!"
	case LoginPasswordNotCorrect: return "The password is not correct!"
	case LoginSuccess: return "Login Successfully!"
	case LoginError: return "Fail to login!"
	default : return "Unkown"
	}
}