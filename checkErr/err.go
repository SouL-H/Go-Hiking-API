package checkerr
//Error control structure.
func CheckError(err error){
	if err != nil{
		panic(err)
	}
}