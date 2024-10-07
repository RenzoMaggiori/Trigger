import React, { useState } from 'react';
import { View, Text, StyleSheet } from 'react-native';
// @ts-ignore
import { PanGestureHandler, PanGestureHandlerGestureEvent } from 'react-native-gesture-handler';
import Animated, { useAnimatedGestureHandler, useAnimatedStyle, useSharedValue } from 'react-native-reanimated';

export default function Draggable({children}:  any) {

    const translationX = useSharedValue(0);
    const translationY = useSharedValue(0);

    const panGesture = useAnimatedGestureHandler<PanGestureHandlerGestureEvent>({
        onStart: (event, ctx) => {
            ctx.x = translationX.value;
            ctx.y = translationY.value;
        },
        onActive: (event, ctx) => {
            translationX.value = ctx.x + event.translationX;
            translationY.value = ctx.y + event.translationY;
        },
    });

    const animatedStyle = useAnimatedStyle(() => {
        return {
            transform: [{ translateX: translationX.value }, { translateY: translationY.value }]
        };
    });

    return (
        <PanGestureHandler onGestureEvent={panGesture}>
            <Animated.View style={animatedStyle}>
                {children}
            </Animated.View>
        </PanGestureHandler>
        
    );
}
