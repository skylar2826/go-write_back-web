package handler_func

import (
	"fmt"
	"geektime-go2/init_register"
	"geektime-go2/orm/orm_gen/data"
	"geektime-go2/web/context"
	"log"
)

//func (u *UserNoSecret) MarshalBinary() ([]byte, error) {
//	var buf bytes.Buffer
//
//	nameLength := uint8(len(u.Name))
//	if err := binary.Write(&buf, binary.BigEndian, nameLength); err != nil {
//		return nil, err
//	}
//	if _, err := buf.WriteString(u.Name); err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//}
//
//func (u *UserNoSecret) UnmarshalBinary(data []byte) error {
//	r := bytes.NewReader(data)
//	var nameLength uint8
//	if err := binary.Read(r, binary.BigEndian, &nameLength); err != nil {
//		return err
//	}
//	nameBytes := make([]byte, nameLength)
//	if _, err := io.ReadFull(r, nameBytes); err != nil {
//		return err
//	}
//	u.Name = string(nameBytes)
//	return nil
//}

func SignUp(c *context.Context) {
	//u := &data.User{}
	//u.Username = c.R.FormValue("Username")
	//u.Email = c.R.FormValue("Email")

	//ctx, tx, err := init_register.DB.BeginTxV2(c, &sql.TxOptions{})
	//if err != nil {
	//	_ = c.SystemErrorJson(err)
	//}
	//
	//s := selector.NewSelector[data.User](tx).Where(data.UserUsernameEq(u.Username), data.UserEmailEq(u.Email))
	//var res any
	//res, err = s.Get(ctx)
	//val, ok := res.(*data.User)

	res, err := init_register.Cache.Get(*c, c.R.FormValue("Id"))
	if err != nil {
		_ = c.SystemErrorJson(err)
	}

	val, ok := res.(*data.User)

	if ok {
		_ = c.OkJson(fmt.Sprintf("200 Id: %d, Birthdate: %s, Email: %s, BaseInfo:%v\n", val.Id, val.Birthdate, val.Email, val.BaseInfo))
	} else {
		err = c.UnauthorizedJsonDirect(c.R.URL.Path)
		if err != nil {
			er := c.SystemErrorJson(err)
			if er != nil {
				log.Fatal("system error: ", er)
			}
		}
	}
}
