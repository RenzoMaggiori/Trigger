import { TriggersService } from '@/api/triggers/service';
import { Colors } from '@/constants/Colors';
import React, { useState } from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { useFocusEffect } from '@react-navigation/native';

interface Node {
    node_id: string;
    action_id: string;
    input: Record<string, any>;
    parents: string[];
    children: string[];
    x_pos: number;
    y_pos: number;
}

interface Template {
    name: string;
    nodes: Node[];
}

export default function TemplatesScreen() {
    const [templates, setTemplates] = useState<Template[]>([]);

    useFocusEffect(
        React.useCallback(() => {
            const loadTemplates = async () => {
                try {
                    const fetchedTemplates = await TriggersService.getTemplates();
                    setTemplates(fetchedTemplates);
                } catch (error) {
                    console.error('Failed to load templates:', error);
                }
            };

            loadTemplates();
        }, [])
    );

    return (
        <ScrollView style={styles.container}>
            <View style={styles.triggerListContainer}>
                <Text style={styles.title}>Available Templates</Text>
                {templates.length > 0 ? (
                    templates.map((template) => (
                        <View key={template.name} style={styles.workspaceCard}>
                            <Text style={styles.workspaceTitle}>{template.name}</Text>
                            {template.nodes.map((node) => {
                                const isTriggerOrHeadNode =
                                    node.parents.length === 0; // Nodo cabeza si no tiene padres

                                return (
                                    <View
                                        key={node.node_id}
                                        style={[
                                            styles.nodeCard,
                                            isTriggerOrHeadNode && styles.triggerNodeCard,
                                        ]}
                                    >
                                        <Text
                                            style={[
                                                styles.nodeText,
                                                isTriggerOrHeadNode && styles.whiteText,
                                            ]}
                                        >
                                            {node.node_id}
                                        </Text>
                                        <Text
                                            style={[
                                                styles.nodeText,
                                                isTriggerOrHeadNode && styles.whiteText,
                                            ]}
                                        >
                                            Action ID: {node.action_id}
                                        </Text>
                                    </View>
                                );
                            })}
                        </View>
                    ))
                ) : (
                    <Text style={styles.noTriggersText}>No templates available.</Text>
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
    triggerListContainer: {
        marginTop: 10,
        marginBottom: 30,
    },
    title: {
        fontSize: 22,
        fontWeight: 'bold',
        marginBottom: 16,
        textAlign: 'center',
    },
    workspaceCard: {
        backgroundColor: Colors.light.tintLight,
        borderRadius: 10,
        marginBottom: 20,
        padding: 16,
        shadowColor: '#000',
        shadowOpacity: 0.1,
        shadowRadius: 5,
        shadowOffset: { width: 0, height: 2 },
    },
    workspaceTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 10,
    },
    nodeCard: {
        backgroundColor: '#f9f9f9',
        borderRadius: 8,
        padding: 10,
        shadowColor: '#000',
        shadowOpacity: 0.05,
        shadowRadius: 3,
        shadowOffset: { width: 0, height: 1 },
        marginBottom: 10,
    },
    triggerNodeCard: {
        backgroundColor: Colors.light.tintDark,
    },
    nodeText: {
        fontSize: 14,
        color: Colors.light.tintDark,
    },
    whiteText: {
        color: '#fff',
    },
    noTriggersText: {
        textAlign: 'center',
        color: '#888',
    },
});

// import { TriggersService } from '@/api/triggers/service';
// import { Colors } from '@/constants/Colors';
// import React, { useState } from 'react';
// import { View, Text, StyleSheet, ScrollView } from 'react-native';
// import { useFocusEffect } from '@react-navigation/native';

// interface Node {
//     node_id: string;
//     action_id: string;
//     input: Record<string, any>;
//     parents: string[];
//     children: string[];
//     x_pos: number;
//     y_pos: number;
// }

// interface Template {
//     name: string;
//     nodes: Node[];
// }

// export default function TemplatesScreen() {
//     const [templates, setTemplates] = useState<Template[]>([]);

//     useFocusEffect(
//         React.useCallback(() => {
//             const loadTemplates = async () => {
//                 try {
//                     const fetchedTemplates = await TriggersService.getTemplates();
//                     setTemplates(fetchedTemplates);
//                 } catch (error) {
//                     console.error('Failed to load templates:', error);
//                 }
//             };

//             loadTemplates();
//         }, [])
//     );

//     return (
//         <ScrollView style={styles.container}>
//             <View style={styles.triggerListContainer}>
//                 <Text style={styles.title}>Available Templates</Text>
//                 {templates.length > 0 ? (
//                     templates.map((template) => (
//                         <View key={template.name} style={styles.workspaceCard}>
//                             <Text style={styles.workspaceTitle}>{template.name}</Text>
//                             {template.nodes.map((node) => (
//                                 <View key={node.node_id} style={styles.nodeCard}>
//                                     <Text style={styles.nodeText}>
//                                         Node ID: {node.node_id}
//                                     </Text>
//                                     <Text style={styles.nodeText}>
//                                         Action ID: {node.action_id}
//                                     </Text>
//                                 </View>
//                             ))}
//                         </View>
//                     ))
//                 ) : (
//                     <Text style={styles.noTriggersText}>No templates available.</Text>
//                 )}
//             </View>
//         </ScrollView>
//     );
// }

// const styles = StyleSheet.create({
//     container: {
//         flex: 1,
//         padding: 16,
//     },
//     triggerListContainer: {
//         marginTop: 10,
//         marginBottom: 30,
//     },
//     title: {
//         fontSize: 22,
//         fontWeight: 'bold',
//         marginBottom: 16,
//         textAlign: 'center',
//     },
//     workspaceCard: {
//         backgroundColor: Colors.light.tintLight,
//         borderRadius: 10,
//         marginBottom: 20,
//         padding: 16,
//         shadowColor: '#000',
//         shadowOpacity: 0.1,
//         shadowRadius: 5,
//         shadowOffset: { width: 0, height: 2 },
//     },
//     workspaceTitle: {
//         fontSize: 18,
//         fontWeight: 'bold',
//         marginBottom: 10,
//     },
//     nodeCard: {
//         backgroundColor: '#f9f9f9',
//         borderRadius: 8,
//         padding: 10,
//         shadowColor: '#000',
//         shadowOpacity: 0.05,
//         shadowRadius: 3,
//         shadowOffset: { width: 0, height: 1 },
//         marginBottom: 10,
//     },
//     nodeText: {
//         fontSize: 14,
//         color: Colors.light.tintDark,
//     },
//     noTriggersText: {
//         textAlign: 'center',
//         color: '#888',
//     },
// });
