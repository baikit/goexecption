package execption

import "strconv"

func ThrowError(err error, msg string, code ...int) {

	if err != nil {

		statusCode := 400

		if len(code) > 0 {
			statusCode = code[0]
		}

		panic("Exception&" + strconv.Itoa(statusCode) + "&" + msg)
	}

}

func ThrowTrue(condition bool, msg string, code ...int) {
	if condition {
		statusCode := 400

		if len(code) > 0 {
			statusCode = code[0]
		}

		panic("Exception&" + strconv.Itoa(statusCode) + "&" + msg)
	}
}
