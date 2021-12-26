package com.coderconsole.tdd;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

class EmailLookUpTest {

    EmailLookUp emailLookUp;

    @BeforeEach
    void setUp() {
        emailLookUp = new EmailLookUp();
    }

    @Test
    void given_valid_email_then_return_true() {
        boolean isValidEmail = emailLookUp.validateEmail("nitesh.tiwari@gmail.com");
        assertThat(isValidEmail).isTrue();

    }

    @Test
    void given_invalid_email_then_return_false() {
        boolean isValidEmail = emailLookUp.validateEmail(null);
        assertThat(isValidEmail).isFalse();
    }

    @Test
    void given_email_id_return_domain_name() {
        String email = emailLookUp.getDomainName("nitesh.tiwari@xyz.com");
        assertThat(email).isEqualTo("xyz.com");
    }

    @Test
    void given_invalid_lookup_email_then_return_false() {
        boolean isValidEmail = emailLookUp.validateEmail("nitesh.tiwari@xyz.com");
        assertThat(isValidEmail).isFalse();
    }


}