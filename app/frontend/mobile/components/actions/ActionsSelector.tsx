import React, { useState, useEffect } from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { Colors } from '@/constants/Colors';
import Button from '@/components/Button';
import { TriggersService } from '@/api/triggers/service';

interface ActionSelectorProps {
    provider: string | null;
    onActionSelect: (action: string) => void;
    type: string;
}

export default function ActionSelector({ provider, onActionSelect, type }: ActionSelectorProps) {
    const [actions, setActions] = useState<any[]>([]);
    const [reactions, setReactions] = useState<any[]>([]);

    useEffect(() => {
        async function fetchActions() {
            const result = await TriggersService.getTriggersByProvider(provider ?? '');
            setActions(result);
        }
        fetchActions();
    }, [provider]);

    useEffect(() => {
        async function fetchReactions() {
            const result = await TriggersService.getReactionsByProvider(provider ?? '');
            setReactions(result);
        }
        fetchReactions();
    }, [provider]);

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Select an Action</Text>
            <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={styles.actionsList}>
                {type === 'trigger' ? (
                    actions.map((action, index) => (
                        <Button
                            key={action.id}
                            onPress={() => onActionSelect(action.action)}
                            title={action.action}
                            textColor={Colors.light.tint}
                            backgroundColor='#fff'
                            borderCol={Colors.light.tint}
                            style={{ marginVertical: 5 }}
                        />
                    ))
                ) : (
                    reactions.map((reaction, index) => (
                        <Button
                            key={reaction.id}
                            onPress={() => onActionSelect(reaction.action)}
                            title={reaction.action}
                            textColor={Colors.light.tint}
                            backgroundColor='#fff'
                            borderCol={Colors.light.tint}
                            style={{ marginVertical: 5 }}
                        />
                    ))
                )}
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
    actionsList: {
        flexDirection: 'column',
    },
});
