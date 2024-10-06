import React from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { Ionicons, FontAwesome5 } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import ButtonIcon from '../ButtonIcon';

interface TechnologySelectorProps {
    onTechSelect: (tech: string) => void;
}

const technologies = [
    { name: 'Google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} /> },
    { name: 'Github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
    { name: 'Outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.outlook} /> },
    { name: 'Slack', icon: <FontAwesome5 name="slack" size={30} color={Colors.light.slack} /> },
    { name: 'Discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.discord} /> },
];

export default function TechnologySelector({ onTechSelect }: TechnologySelectorProps) {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>Select a Service</Text>
            <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={styles.techList}>
                {technologies.map((tech, index) => (
                    <ButtonIcon
                        key={index}
                        title={tech.name}
                        icon={tech.icon}
                        onPress={() => onTechSelect(tech.name)}
                        textColor={Colors.light.tint}
                        borderCol={Colors.light.tint}
                        style={{ marginVertical: 5 }}
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
