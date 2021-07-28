/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    metadef
 * @Date:    2021/6/10 11:14 上午
 * @package: pbc
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package pb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// ====== MsgID ======

func (x MsgID) ID() int32 {
	return int32(x)
}

func (x MsgID) Length() int {
	return 4
}

func (x MsgID) Encode() ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func DecodeMsgID(b []byte) (MsgID, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return MsgID(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return MsgID(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return MsgID(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}

}

// ======== error ========

func (x ErrCode) Error() string {
	return fmt.Sprintf("ErrCode=%d ErrMsg=%s", x, ErrCode_name[int32(x)])
}

func (x ErrCode) ErrMsg() string {
	return ErrCode_name[int32(x)]
}

func (x ErrCode) Code() int32 {
	return int32(x)
}

func (x ErrCode) ECode() ErrCode {
	return x
}
