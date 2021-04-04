package com.coderconsole.tdd;

import java.util.ArrayList;
import java.util.List;

public class PaymentMode {

    final List<String> paymentModes = new ArrayList<>();

    public PaymentMode(){
        paymentModes.add("paytm");
        paymentModes.add("payu");
        paymentModes.add("cc");
    }
    public boolean triggerPayment(String mode) {
        return paymentModes.contains(mode);
    }
}
