package com.coderconsole.tdd.auth;

public class LoginFlow {

    public int isValidMobileNumber(String mobile) {
        if (mobile.length() < 10 || !(mobile.startsWith("6")
                || mobile.startsWith("7") || mobile.startsWith("8")
                || mobile.startsWith("9"))
        ) return 400;

        return 200;
    }
}
