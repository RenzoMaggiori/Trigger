import React from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { Colors } from '@/constants/Colors';
import Button from '@/components/Button';

interface ActionSelectorProps {
    onActionSelect: (id: number) => void;
}

const actions = [
    { id: 1, input: [], output: [], provider:'gmail', type: 'trigger', action: 'new email_1'},
    { id: 2, input: [], output: [], provider:'gmail', type: 'trigger', action: 'new email_2'},
    { id: 3, input: [], output: [], provider:'gmail', type: 'trigger', action: 'new email_3'},
    { id: 4, input: [], output: [], provider:'gmail', type: 'reaction', action: 'send email_1'},
    { id: 5, input: [], output: [], provider:'gmail', type: 'reaction', action: 'send email_2'},
    { id: 6, input: [], output: [], provider:'gmail', type: 'reaction', action: 'send email_3'},
];

export default function ActionSelector({ onActionSelect }: ActionSelectorProps) {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>Select a Provider</Text>
            <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={styles.techList}>
                {actions.map((action, index) => (
                    <Button
                        onPress={() => onActionSelect(action.id)}
                        title={action.action}
                    />
                ))}
            </ScrollView>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        padding: 20,
        backgroundColor: '#fff',
        borderRadius: 10,
        marginHorizontal: 30,
    },
    title: {
        fontSize: 18,
        color: Colors.light.tint,
        fontWeight: 'bold',
        marginBottom: 15,
        textAlign: 'center',
    },
    techList: {
        flexDirection: 'column',
    },
});
