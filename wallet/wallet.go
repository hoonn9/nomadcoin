package wallet

/*
	sign, verify


	1) hash the msg
	"message" -> hash(x) -> "hashed_message"

	2) generate key pair
	private key, public key (using go, save private key to a file)

	3) sign the hash
	("hashed_message" + private key) -> "signature"

	4) verify
	("hashed_message" + "signature" + public key) -> true | false
*/
