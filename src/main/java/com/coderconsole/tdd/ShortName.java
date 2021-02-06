package com.coderconsole.tdd;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class ShortName {
    public String concatInitials(String firstName, String lastName) {

        List<String> list = new ArrayList<>();
        list.add(firstName);
        list.add(lastName);

        Stream<String> objectStream = list.stream()
                .filter(it -> it != null && !it.isEmpty())
                .map(s -> s.substring(0, 1).trim());
        return objectStream.collect(Collectors.joining());

    }

    /*
    public String concatInitialsRefactored2(String firstName, String lastName) {

        String result = "";

        if (firstName != null && !firstName.isEmpty()) {
            result = firstName.substring(0, 1);
        }

        if (lastName != null && !lastName.isEmpty()) {
            result += lastName.substring(0, 1);
        }
        return result;
    }

    public String concatInitialsRefactor1(String firstName, String lastName) {

        if (firstName.isEmpty() && lastName.isEmpty())
            return "";
        else if (!firstName.isEmpty() && !lastName.isEmpty()) {
            return firstName.substring(0, 1) + lastName.substring(0, 1);
        } else if (firstName.isEmpty() && !lastName.isEmpty()) {
            return lastName.substring(0, 1);
        } else {
            return firstName.substring(0, 1);
        }
    }

    */


}
