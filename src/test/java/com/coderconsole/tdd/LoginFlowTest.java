package com.coderconsole.tdd;

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
    void given_mobile_number_return_200() {
        assertThat(loginFlow
                .validDateMobile("9876543210"))
                .isEqualTo(200);
    }

    @Test
    void given_empty_mobile_number_return_400() {
        assertThat(loginFlow
                .validDateMobile(""))
                .isEqualTo(400);
    }

    @Test
    void given_invalid_mobile_number_return_400() {
        assertThat(loginFlow
                .validDateMobile("1"))
                .isEqualTo(400);
    }

    @Test
    void given_invalid_sequence_mobile_number_return_400() {
        assertThat(loginFlow
                .validDateMobile("1111111111"))
                .isEqualTo(400);
    }
}
