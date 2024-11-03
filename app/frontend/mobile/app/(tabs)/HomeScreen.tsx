import React, { useEffect, useState, useCallback } from 'react';
import { View, StyleSheet, ScrollView, Text, TouchableOpacity } from 'react-native';
import ButtonIcon from '@/components/ButtonIcon';
import Button from '@/components/Button';
import { MaterialCommunityIcons, MaterialIcons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import { TriggersService } from '@/api/triggers/service';
import { useFocusEffect } from '@react-navigation/native';

export default function HomeScreen() {
    const [workspaces, setWorkspaces] = useState<Workspace[]>([]);

    useFocusEffect(
        useCallback(() => {
            const loadWorkspaces = async () => {
                try {
                    const fetchedWorkspaces = await TriggersService.getTriggers();
                    // console.log('fetchedWorkspaces:', fetchedWorkspaces);

                    const updatedWorkspaces = await Promise.all(
                        fetchedWorkspaces.map(async (workspace: Workspace) => {
                            const nodesWithDetails = await Promise.all(
                                workspace.nodes.map(async (node) => {
                                    const actionDetails = await TriggersService.getActionById(node.action_id);
                                    return {
                                        ...node,
                                        actionDetails,
                                    };
                                })
                            );
                            return {
                                ...workspace,
                                nodes: nodesWithDetails,
                            };
                        })
                    );

                    setWorkspaces(updatedWorkspaces);
                } catch (error) {
                    console.error('Failed to load workspaces:', error);
                }
            };

            loadWorkspaces();
        }, [])
    );

    return (
        <ScrollView style={styles.container}>
            <PromoItem />
            <Button
                title='Templates'
                onPress={() => console.log('Templates')}
                backgroundColor={Colors.light.tint}
                textColor='#FFFFFF'
                buttonWidth="35%"
                paddingV={7.5}
            />
            <TriggerList workspaces={workspaces} />
        </ScrollView>
    );
}

function PromoItem() {
    return (
        <View style={styles.promoContainer}>
            <View style={styles.promoBox}>
                <Text style={styles.promoText}>Try Trigger for 30 days free</Text>
                <ButtonIcon
                    onPress={() => console.log('Start free trial')}
                    title="Start free trial"
                    icon={<MaterialCommunityIcons name="star-shooting" size={24} color={Colors.light.tint} />}
                    backgroundColor="#FFFFFF"
                    textColor={Colors.light.tint}
                />
            </View>
        </View>
    );
}

interface Workspace {
    id: string;
    user_id: string;
    nodes: {
        node_id: string;
        action_id: string;
        status: string;
        actionDetails: {
            id: string;
            input: string[];
            output: string[];
            provider: string;
            type: string;
            action: string;
        };
    }[];
}

// function TriggerList({ workspaces }: { workspaces: Workspace[] }) {
//     return (
//         <View style={styles.triggerListContainer}>
//             <Text style={styles.title}>Your Workspaces</Text>
//             {workspaces.length > 0 ? (
//                 workspaces.map((workspace) => (
//                     <View key={workspace.id} style={styles.workspaceCard}>
//                         <Text style={styles.workspaceTitle}>Workspace ID: {workspace.id}</Text>
//                         <View style={styles.nodesContainer}>
//                             {workspace.nodes.map((node) => {
//                                 const isTrigger = node.actionDetails?.type === 'trigger';

//                                 return (
//                                     <View
//                                         key={node.node_id}
//                                         style={[
//                                             styles.nodeCard,
//                                             isTrigger && styles.triggerNodeCard,
//                                         ]}
//                                     >
//                                         <View style={styles.nodeHeader}>
//                                             <View style={styles.nodeDetails}>
//                                                 <Text
//                                                     style={[
//                                                         styles.nodeText,
//                                                         isTrigger && styles.whiteText,
//                                                     ]}
//                                                 >
//                                                     Provider: {node.actionDetails?.provider}
//                                                 </Text>
//                                                 <Text
//                                                     style={[
//                                                         styles.nodeText,
//                                                         isTrigger && styles.whiteText,
//                                                     ]}
//                                                 >
//                                                     {isTrigger ? 'Action' : 'Reaction'}: {node.actionDetails?.action}
//                                                 </Text>
//                                                 <View style={styles.actionDetailsContainer}>
//                                                     <Text
//                                                         style={[
//                                                             styles.actionDetailText,
//                                                             isTrigger && styles.whiteText,
//                                                         ]}
//                                                     >
//                                                         Inputs: {node.actionDetails?.input.join(', ')}
//                                                     </Text>
//                                                     <Text
//                                                         style={[
//                                                             styles.actionDetailText,
//                                                             isTrigger && styles.whiteText,
//                                                         ]}
//                                                     >
//                                                         Outputs: {node.actionDetails?.output.join(', ')}
//                                                     </Text>
//                                                 </View>
//                                             </View>
//                                             <View style={styles.rightSection}>
//                                                 <Text
//                                                     style={[
//                                                         styles.nodeStatus,
//                                                         isTrigger && styles.whiteText,
//                                                     ]}
//                                                 >
//                                                     {node.status === 'active' ? 'Active' : 'Inactive'}
//                                                 </Text>
//                                                 <TouchableOpacity
//                                                     style={[
//                                                         styles.actionButton,
//                                                         isTrigger && styles.invertedActionButton,
//                                                     ]}
//                                                     onPress={() =>
//                                                         console.log(`${node.status === 'active' ? 'Stop' : 'Start'} action`)
//                                                     }
//                                                 >
//                                                     <MaterialIcons
//                                                         name={node.status === 'active' ? 'pause' : 'play-arrow'}
//                                                         size={16}
//                                                         color={isTrigger ? Colors.light.tintDark : '#fff'}
//                                                     />
//                                                     <Text
//                                                         style={[
//                                                             styles.actionButtonTxt,
//                                                             isTrigger && styles.invertedActionButtonTxt,
//                                                         ]}
//                                                     >
//                                                         {node.status === 'active' ? 'Stop' : 'Start'}
//                                                     </Text>
//                                                 </TouchableOpacity>
//                                             </View>
//                                         </View>
//                                     </View>
//                                 );
//                             })}
//                         </View>
//                     </View>
//                 ))
//             ) : (
//                 <Text style={styles.noTriggersText}>No workspaces available.</Text>
//             )}
//         </View>
//     );
// }

function TriggerList({ workspaces }: { workspaces: Workspace[] }) {
    const handleAction = async (node: any) => {
        try {
            if (node.status === 'active') {
                await TriggersService.stopAction(node.action_id);
                console.log(`Stopped action with ID: ${node.action_id}`);
            } else {
                await TriggersService.startAction(node.action_id);
                console.log(`Started action with ID: ${node.action_id}`);
            }
        } catch (error) {
            console.error("Error handling action:", error);
        }
    };

    return (
        <View style={styles.triggerListContainer}>
            <Text style={styles.title}>Your Workspaces</Text>
            {workspaces.length > 0 ? (
                workspaces.map((workspace) => (
                    <View key={workspace.id} style={styles.workspaceCard}>
                        <Text style={styles.workspaceTitle}>Workspace ID: {workspace.id}</Text>
                        <View style={styles.nodesContainer}>
                            {workspace.nodes.map((node) => {
                                const isTrigger = node.actionDetails?.type === 'trigger';

                                return (
                                    <View
                                        key={node.node_id}
                                        style={[
                                            styles.nodeCard,
                                            isTrigger && styles.triggerNodeCard,
                                        ]}
                                    >
                                        <View style={styles.nodeHeader}>
                                            <View style={styles.nodeDetails}>
                                                <Text
                                                    style={[
                                                        styles.nodeText,
                                                        isTrigger && styles.whiteText,
                                                    ]}
                                                >
                                                    Provider: {node.actionDetails?.provider}
                                                </Text>
                                                <Text
                                                    style={[
                                                        styles.nodeText,
                                                        isTrigger && styles.whiteText,
                                                    ]}
                                                >
                                                    {isTrigger ? 'Action' : 'Reaction'}: {node.actionDetails?.action}
                                                </Text>
                                                <View style={styles.actionDetailsContainer}>
                                                    <Text
                                                        style={[
                                                            styles.actionDetailText,
                                                            isTrigger && styles.whiteText,
                                                        ]}
                                                    >
                                                        Inputs: {node.actionDetails?.input.join(', ')}
                                                    </Text>
                                                    <Text
                                                        style={[
                                                            styles.actionDetailText,
                                                            isTrigger && styles.whiteText,
                                                        ]}
                                                    >
                                                        Outputs: {node.actionDetails?.output.join(', ')}
                                                    </Text>
                                                </View>
                                            </View>
                                            <View style={styles.rightSection}>
                                                <Text
                                                    style={[
                                                        styles.nodeStatus,
                                                        isTrigger && styles.whiteText,
                                                    ]}
                                                >
                                                    {node.status === 'active' ? 'Active' : 'Inactive'}
                                                </Text>
                                                <TouchableOpacity
                                                    style={[
                                                        styles.actionButton,
                                                        isTrigger && styles.invertedActionButton,
                                                    ]}
                                                    onPress={() => handleAction(node)}
                                                >
                                                    <MaterialIcons
                                                        name={node.status === 'active' ? 'pause' : 'play-arrow'}
                                                        size={16}
                                                        color={isTrigger ? Colors.light.tintDark : '#fff'}
                                                    />
                                                    <Text
                                                        style={[
                                                            styles.actionButtonTxt,
                                                            isTrigger && styles.invertedActionButtonTxt,
                                                        ]}
                                                    >
                                                        {node.status === 'active' ? 'Stop' : 'Start'}
                                                    </Text>
                                                </TouchableOpacity>
                                            </View>
                                        </View>
                                    </View>
                                );
                            })}
                        </View>
                    </View>
                ))
            ) : (
                <Text style={styles.noTriggersText}>No workspaces available.</Text>
            )}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
    },
    promoContainer: {
        alignItems: 'center',
        marginBottom: 10,
    },
    promoBox: {
        width: '100%',
        backgroundColor: Colors.light.tint,
        padding: 10,
        borderRadius: 10,
        alignItems: 'center',
    },
    promoText: {
        color: '#fff',
        fontSize: 18,
        marginBottom: 10,
        textAlign: 'center',
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
    nodesContainer: {
        marginBottom: 10,
    },
    nodeCard: {
        backgroundColor: '#f9f9f9',
        borderRadius: 8,
        padding: 10,
        marginBottom: 10,
        shadowColor: '#000',
        shadowOpacity: 0.05,
        shadowRadius: 3,
        shadowOffset: { width: 0, height: 1 },
    },
    triggerNodeCard: {
        backgroundColor: Colors.light.tintDark,
    },
    nodeHeader: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'flex-start',
    },
    nodeDetails: {
        flex: 1,
    },
    rightSection: {
        flexDirection: 'column',
        alignItems: 'flex-end',
    },
    nodeText: {
        fontSize: 14,
        fontWeight: 'bold',
        color: Colors.light.tintDark,
    },
    whiteText: {
        color: '#fff',
    },
    actionDetailsContainer: {
        marginTop: 5,
    },
    actionDetailText: {
        fontSize: 12,
    },
    nodeStatus: {
        fontSize: 14,
        fontWeight: 'bold',
        color: '#666',
        marginBottom: 5,
    },
    actionButton: {
        flexDirection: 'row',
        alignItems: 'center',
        backgroundColor: Colors.light.tint,
        paddingVertical: 4,
        paddingHorizontal: 8,
        borderRadius: 20,
        marginTop: 5,
    },
    invertedActionButton: {
        backgroundColor: '#fff',
    },
    actionButtonTxt: {
        color: '#fff',
        marginLeft: 5,
    },
    invertedActionButtonTxt: {
        color: Colors.light.tintDark,
    },
    noTriggersText: {
        textAlign: 'center',
        color: '#888',
    },
});
