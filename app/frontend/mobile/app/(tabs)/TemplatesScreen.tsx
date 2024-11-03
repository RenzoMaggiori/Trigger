import { TriggersService } from '@/api/triggers/service';
import { Colors } from '@/constants/Colors';
import React, { useEffect, useState } from 'react';
import { View, Text, StyleSheet, ScrollView, TouchableOpacity } from 'react-native';

interface Template {
    id: number;
    provider: string;
    type: string;
    action: string;
}

export default function TemplatesScreen() {
    const [templates, setTemplates] = useState<Template[]>([]);

    useEffect(() => {
        const loadTemplates = async () => {
            const fetchedTemplates = await TriggersService.getTemplates();
            setTemplates(fetchedTemplates);
        };

        loadTemplates();
    }, []);

    return (
        <ScrollView style={styles.container}>
            <View style={styles.cardContainer}>
                {templates.length > 0 ? (
                    templates.map((template) => (
                        <View key={template.id} style={styles.card}>
                            <Text style={styles.title}>{template.provider}</Text>
                            <Text style={styles.textInfo}>{template.type}: {template.action}</Text>
                        </View>
                    ))
                ) : (
                    <Text>No templates available.</Text>
                )}
            </View>
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
    },
    cardContainer: {
        marginBottom: 40,
    },
    card: {
        backgroundColor: Colors.light.tintDark,
        borderRadius: 10,
        padding: 16,
        marginBottom: 10,
        shadowColor: '#000',
        shadowOpacity: 0.1,
        shadowRadius: 5,
        shadowOffset: { width: 0, height: 2 },
        color: '#FFFFFF',
    },
    title: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 5,
        color: '#FFFFFF',
    },
    textInfo: {
        color: '#FFFFFF',
    }
});