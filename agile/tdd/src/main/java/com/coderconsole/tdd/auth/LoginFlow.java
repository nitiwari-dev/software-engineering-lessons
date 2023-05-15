package com.coderconsole.tdd.auth;

public class LoginFlow {

    public int isValidMobileNumber(String mobile) {
        if (mobile.length() < 10 || !(mobile.startsWith("6")
                || mobile.startsWith("7") || mobile.startsWith("8")
                || mobile.startsWith("9"))
        ) return HTTP_INVALID_INPUT_400;

        return HTTP_OK_200;
    }

    private static final int HTTP_OK_200 = 200;
    private static final int HTTP_INVALID_INPUT_400 = 400;
}
