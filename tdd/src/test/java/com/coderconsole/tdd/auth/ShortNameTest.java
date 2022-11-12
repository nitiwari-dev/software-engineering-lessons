package com.coderconsole.tdd.auth;

import com.coderconsole.tdd.auth.ShortName;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

class ShortNameTest {


    private ShortName shortName;

    @BeforeEach
    public void setUp() {
        shortName = new ShortName();
    }


    @Test
    @DisplayName("Given firstName and last Name concat initial characters")
    void shorName() {
        String result = shortName.concatInitials("Nitesh", "Tiwari");
        assertThat(result).isEqualTo("NT");
    }

    @Test
    @DisplayName("Given empty or null firstName return only last name initial character")
    void firstNameEmpty() {
        String result = shortName.concatInitials("", "Tiwari");
        assertThat(result).isEqualTo("T");
    }

    @Test
    @DisplayName("Given empty or null lastName return only first name initial character")
    void lastNameEmpty() {
        String result = shortName.concatInitials("Nitesh", "");
        assertThat(result).isEqualTo("N");
    }

    @Test
    @DisplayName("Given empty firstName and lastName return empty")
    void emptyFirstNameAndLastName() {
        String result = shortName.concatInitials("", "");
        assertThat(result).isEmpty();
    }


    @Test
    @DisplayName("Given spaced firstName with lastName return trim firstname return lastName initials")
    void spacedFirstNameWithLastName() {
        String result = shortName.concatInitials("  ", "");
        assertThat(result).isEmpty();
    }
}