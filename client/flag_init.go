package main

import "flag"

func flag_init() (user_name, addr string, err error) {
	flag.StringVar(&user_name, "n", "", "User name")
	flag.StringVar(&addr, "t", ":8080", "Address of the server")
	flag.Parse()
	if user_name == "" {
		return user_name, addr, &NoUserNameError{}
	} else if flag.NArg() > 0 {
		return user_name, addr, &BadArgNumError{}
	} else if len(user_name) > 256 {
		return user_name, addr, &TooLongNameError{}
	}
	return user_name, addr, nil
}

type NoUserNameError struct {
}

func (err *NoUserNameError) Error() string {
	return "No user name"
}

type TooLongNameError struct {
}

func (err *TooLongNameError) Error() string {
	return "User name is too long"
}

type BadArgNumError struct {
}

func (err *BadArgNumError) Error() string {
	return "Bad number of arguments"
}
