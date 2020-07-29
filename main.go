//
// If your payment channel create message receipt looks like this:
//
//	{ exit_code: 0, ret: "gkMA8TdVAuEnBlturfm30tDuzwd59jqHcDfr", gas_used: 8897444 }
//
// extract the payment channel address like this:
//
// echo "gkMA8TdVAuEnBlturfm30tDuzwd59jqHcDfr" | openssl enc -d -a | ./pchmarshaller
//

//
// Compile:  `go build -o pchmarshaller main.go`
//

package main

import(
	"fmt"
	"bytes"
	"bufio"
	"os"
	"io"

	init_ "github.com/filecoin-project/specs-actors/actors/builtin/init"
)

//type MessageReceipt struct {
//	ExitCode int64
//	Return   []byte
//	GasUsed  int64
//}

func main() {
	var ret []byte
	in := bufio.NewReader(os.Stdin)
	for {
		c, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		ret = append(ret, c)
	}

	//mwait := MessageReceipt{
	//	ExitCode: 0,
	//	Return:  []byte{0x82,0x43,0x00,0xF1,0x37,0x55,0x02,0xE1,0x27,0x06,0x5B,0x6E,0xAD,0xF9,0xB7,0xD2,0xD0,0xEE,0xCF,0x07,0x79,0xF6,0x3A,0x87,0x70,0x37,0xEB},
	//	GasUsed: 8897444,
	//}

	var decodedReturn init_.ExecReturn
	err := decodedReturn.UnmarshalCBOR(bytes.NewReader(ret))
	if err != nil {
		fmt.Printf("err=%v",err)
		return
	}
	paychaddr := decodedReturn.RobustAddress
	fmt.Printf("paychaddr = '%s'\n",paychaddr);
}
