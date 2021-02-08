package com.coderconsole.tdd;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class ShortName {
    public String concatInitials(String firstName, String lastName) {

        List<String> list = new ArrayList<>();
        list.add(firstName);
        list.add(lastName);

        return list.stream()
                .filter(this::isValidShortName)
                .map(s -> String.valueOf(s.charAt(0))).collect(Collectors.joining());
    }

    private boolean isValidShortName(String it) {
        return it != null && !it.trim().isEmpty();
    }

    /*
     public String concatInitials(String firstName, String lastName) {

        StringBuilder result = new StringBuilder();

        if (firstName != null && !firstName.trim().isEmpty()) {
            result.append(firstName.charAt(0));
        }

        if (lastName != null && !lastName.trim().isEmpty()) {
            result.append(lastName.charAt(0));
        }
        return result.toString();
    }

    public String concatInitials(String firstName, String lastName) {

        String result = "";

        if (firstName != null && !firstName.trim().isEmpty()) {
            result = firstName.substring(0, 1);
        }

        if (lastName != null && !lastName.trim().isEmpty()) {
            result += lastName.substring(0, 1);
        }
        return result;
    }

    public String concatInitials(String firstName, String lastName) {

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
