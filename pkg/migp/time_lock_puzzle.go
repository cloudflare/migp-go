package main
import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"

	"fmt"
	"io"
	"math/big"
	"time"
)
type TimeLockPuzzleConfig struct {
	N 			*big.Int
	phi_n 		*big.Int
	T 			*big.Int
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}


func DefaultTimeLockPuzzleConfig() TimeLockPuzzleConfig {

	private_key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	primes := private_key.Primes
	N := private_key.D
	p_minus_1  := new(big.Int).Sub(primes[0], big.NewInt(1))
	q_minus_1 := new(big.Int).Sub(primes[1], big.NewInt(1))
	phi_n := new(big.Int).Mul(p_minus_1 , q_minus_1)

	return TimeLockPuzzleConfig {
		N: N,
		phi_n: phi_n,
		T: big.NewInt(2048), // todo: how to set t
	}
}

func time_lock_puzzle(message []byte, timeLockPuzzleConfig TimeLockPuzzleConfig) ([]byte, *big.Int, *big.Int) {


	// creating a ephemeral_key to lock the message
	ephemeral_key := make([]byte, 32)
	rand.Read(ephemeral_key)

	

	//fmt.Println(ephemeral_key)
	// creating a new block cipher with the generated `ephemeral_key`.
	c, err := aes.NewCipher(ephemeral_key)
	if err != nil {
        fmt.Println(err)
    }
	// wrapping the block cipher in Galois Counter Mode (GCM)
	gcm, err := cipher.NewGCM(c) 
	if err != nil {
        fmt.Println(err)
    }
	// creating nonce for the symmetric encryption.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err)
    }
	// doing encryption
	c_m := gcm.Seal(nonce, nonce, message, nil)
	
	// generating a
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
    	panic(err)
	}
	
	a := new(big.Int)
	a.SetBytes(token)
	// 1 < a < n
	a.Mod(a, timeLockPuzzleConfig.N)
	a.Add(a, big.NewInt(1))
	
	K := new(big.Int)
	K.SetBytes(ephemeral_key)
	K.Mod(K, timeLockPuzzleConfig.N)

	
	
	// e = 2^t mod \phi(n)
	e := new(big.Int).Exp(big.NewInt(2), timeLockPuzzleConfig.T, timeLockPuzzleConfig.phi_n) // e = 2^t mod \phi(n)	
	// b  = a^e mod n
	b :=  new(big.Int).Exp(a, e, timeLockPuzzleConfig.N)
	// c_k = K + b) mod n
	c_k := new(big.Int).Add(K, b)
	return c_m, c_k, a
}

func time_unlock_puzzle(c_m []byte, c_k *big.Int , T *big.Int , a *big.Int , N *big.Int  ) []byte {
	
	
	start := makeTimestamp()
	b := a
	p := new(big.Int).Exp(big.NewInt(2), T, nil) // p = 2^T
	b.Exp(b, p , N)
	//fmt.Println("Size before sub", len(b.Bytes())) 
	end := makeTimestamp()
	fmt.Println("Time taken for squrings = ", end-start)
	
	
	K := new(big.Int).Sub(c_k, b) // this line should decrease the size of K to 32 bytes.
	
	for K.Cmp(big.NewInt(0)) < 0 {
			K.Add(K, N)
	}
	K.Mod(K, N)
	//fmt.Println("Size after sub", len(K.Bytes()))
	//fmt.Println(K)

	ephemeral_key := K.Bytes()
	fmt.Println("Size ", len(ephemeral_key))
	c, err := aes.NewCipher(ephemeral_key)
	if err != nil {
        fmt.Println(err)
    }

	gcm, err := cipher.NewGCM(c) // c is from key
	if err != nil {
        fmt.Println(err)
    }
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := c_m[:nonceSize], c_m[nonceSize:]
	decryptedPlainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
        fmt.Println(err)
    }
	return decryptedPlainText
}



func main() {

	// MIGP server side: creating configs from default configuration
	var timeLockPuzzleConfig TimeLockPuzzleConfig
	timeLockPuzzleConfig = DefaultTimeLockPuzzleConfig()
	
	// MIGP server side: create a timelock puzzle <n a t c_k, c_m> for the client. ref.  https://people.csail.mit.edu/rivest/pubs/RSW96.pdf section 2.1
	message := []byte("This is a test message")
	start1 := makeTimestamp()
	c_m, c_k, a := time_lock_puzzle(message, timeLockPuzzleConfig)
	end1 := makeTimestamp()
	fmt.Println("Time to lock the puzzle", end1-start1)
	fmt.Println(c_m, c_k, timeLockPuzzleConfig.T, a, timeLockPuzzleConfig.N)
	// client side: unlocking the puzzle with <n a t c_k, c_m> given.
	fmt.Println("Unlocking the puzzle on the client side...")
	start2 := makeTimestamp()
	recovered_message := time_unlock_puzzle(c_m, c_k, timeLockPuzzleConfig.T, a, timeLockPuzzleConfig.N)
	end2 := makeTimestamp()
	fmt.Println("Time to unlock the puzzle", end2-start2)
	fmt.Println(recovered_message)
}