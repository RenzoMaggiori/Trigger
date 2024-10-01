import React from 'react';
import { TouchableOpacity, Text, View, StyleSheet, ViewStyle } from 'react-native';

interface ButtonIconProps {
    onPress: () => void;
    title: string;
    icon: React.ReactElement;
    backgroundColor?: string;
    textColor?: string;
    borderCol?: string;
    style?: ViewStyle;
}

const ButtonIcon: React.FC<ButtonIconProps> = ({ onPress, title, icon, backgroundColor = '#FFFFFF', textColor = '#000000', borderCol = '#ddd', style }) => {
    return (
        <TouchableOpacity style={[styles.button, { backgroundColor, borderColor: borderCol }, style]} onPress={onPress}>
            <View style={styles.contentContainer}>
                {icon}
                <Text style={[styles.buttonText, { color: textColor }]}>{title}</Text>
            </View>
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    button: {
        borderWidth: 1,
        padding: 15,
        borderRadius: 30,
        alignItems: 'center',
        justifyContent: 'center',
        marginVertical: 5,
    },
    contentContainer: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    buttonText: {
        fontSize: 16,
        marginLeft: 10,
    },
});

export default ButtonIcon;
