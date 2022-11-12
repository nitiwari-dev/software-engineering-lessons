package com.coderconsole.tdd.checkout;

import com.coderconsole.tdd.checkout.ShoppingCart;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

import java.util.ArrayList;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.BDDMockito.given;
import static org.mockito.Mockito.verify;

//Write the test
//make is fail
//write the code, make is pass
//go to step 1
class ShoppingCartTest {

    ShoppingCart shoppingCart;

    @BeforeEach
    void setUp(){
        shoppingCart = Mockito.mock(ShoppingCart.class);
    }

    @Test
    void givenShoppingCartIsEmptyReturnEmptyList(){
        List<ShoppingCart.Cart> cartList = shoppingCart.getCartList();
        assertThat(cartList).isEmpty();
    }

    @Test
    void givenShoppingCartHasCartReturnList(){
        given(shoppingCart.getCartList()).willReturn(getMockList());
        List<ShoppingCart.Cart> cartList = shoppingCart.getCartList();
        assertThat(cartList.size()).isEqualTo(1);
        verify(shoppingCart).getCartList();
    }

    @Test
    void givenShoppingCartHasAddCardReturnList(){
        given(shoppingCart.addCart(anyString(), anyLong())).willReturn(getMockList());
        List<ShoppingCart.Cart> cartList = shoppingCart.addCart("A", 29);
        assertThat(cartList.size()).isEqualTo(1);
        verify(shoppingCart).addCart("A", 29);
    }

    List<ShoppingCart.Cart> getMockList() {
        List<ShoppingCart.Cart> carts = new ArrayList<>();
        carts.add(new ShoppingCart.Cart("Onion", 30));
        return carts;
    }


}