package com.coderconsole.tdd.auth;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

class LoginFlowTest {

    LoginFlow loginFlow;

    @BeforeEach
    void setUp() {
        loginFlow = new LoginFlow();
    }

    @Test
    void given_valid_sequence_mobile_number_return_200() {
        assertThat(loginFlow
                .isValidMobileNumber("9876543210"))
                .isEqualTo(200);
    }

    @Test
    void given_empty_mobile_number_return_400() {
        assertThat(loginFlow
                .isValidMobileNumber(""))
                .isEqualTo(400);
    }

    @Test
    void given_invalid_length_mobile_number_return_400() {
        assertThat(loginFlow
                .isValidMobileNumber("1"))
                .isEqualTo(400);
    }

    @Test
    void given_invalid_sequence_mobile_number_return_400() {
        assertThat(loginFlow
                .isValidMobileNumber("1111111111"))
                .isEqualTo(400);
    }
}
