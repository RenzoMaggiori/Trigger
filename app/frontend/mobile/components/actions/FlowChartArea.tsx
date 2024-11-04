import React, { useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, ScrollView, TextInput } from 'react-native';
import { MaterialIcons, Entypo } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';

interface FlowChartAreaProps {
    flow: {
        provider: string,
        action: { id: string, name: string },
        reactions: { id: string, provider: string, name: string }[]
    }[];
    onAddReaction: (actionIndex: number) => void;
    onRemoveAction: (actionIndex: number) => void;
    onSaveTrigger: (actionIndex: number, inputValues: { [key: string]: string }) => void;
}

export default function FlowChartArea({ flow, onAddReaction, onRemoveAction, onSaveTrigger }: FlowChartAreaProps) {
    const [selectedItem, setSelectedItem] = useState<{ type: 'action' | 'reaction', actionIndex: number | null, reactionIndex?: number | null }>({
        type: 'action',
        actionIndex: null,
        reactionIndex: null,
    });

    const [inputValues, setInputValues] = useState<{ [key: string]: string }>({});

    const handleSelectAction = (actionIndex: number) => {
        setSelectedItem(prevState => ({
            type: 'action',
            actionIndex: prevState.actionIndex === actionIndex ? null : actionIndex,
            reactionIndex: null,
        }));
    };

    const handleSelectReaction = (actionIndex: number, reactionIndex: number) => {
        setSelectedItem(prevState => ({
            type: 'reaction',
            actionIndex,
            reactionIndex: prevState.reactionIndex === reactionIndex ? null : reactionIndex,
        }));
    };

    const handleInputChange = (key: string, value: string) => {
        setInputValues(prevValues => ({
            ...prevValues,
            [key]: value,
        }));
    };

    const handleSaveTrigger = (actionIndex: number) => {
        onSaveTrigger(actionIndex, inputValues);
    };

    const renderInputs = (type: 'action' | 'reaction', provider: string, index: number) => {
        const baseKey = `${type}-${provider}-${index}`;

        if (type === 'action') {
            switch (provider) {
                case 'gmail':
                    return (
                        <>
                            <Text style={styles.inputLabel}>From:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="example@example.com"
                                value={inputValues[`${baseKey}-destination`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-destination`, text)}
                            />
                        </>
                    );
                case 'github':
                    return (
                        <>
                            <Text style={styles.inputLabel}>Repository Owner:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter owner name"
                                value={inputValues[`${baseKey}-owner`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-owner`, text)}
                            />
                            <Text style={styles.inputLabel}>Repository Name:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter repository name"
                                value={inputValues[`${baseKey}-repo`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-repo`, text)}
                            />
                        </>
                    );
                case 'twitch':
                    return (
                        <>
                            <Text style={styles.inputLabel}>Channel Name:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter channel name"
                                value={inputValues[`${baseKey}-channel`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-channel`, text)}
                            />
                        </>
                    );
                case 'spotify':
                    return (
                        <>
                            <Text style={styles.inputLabel}>Playlist ID:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter playlist ID"
                                value={inputValues[`${baseKey}-playlistId`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-playlistId`, text)}
                            />
                        </>
                    );
                default:
                    return <Text style={styles.noInputsText}>No specific inputs for this provider.</Text>;
            }
        } else if (type === 'reaction') {
            switch (provider) {
                case 'Slack':
                    return (
                        <>
                            <Text style={styles.inputLabel}>Channel ID:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter channel ID"
                                value={inputValues[`${baseKey}-channelId`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-channelId`, text)}
                            />
                            <Text style={styles.inputLabel}>Message:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter message"
                                value={inputValues[`${baseKey}-message`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-message`, text)}
                            />
                        </>
                    );
                case 'Discord':
                    return (
                        <>
                            <Text style={styles.inputLabel}>Channel ID:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="1234567890"
                                value={inputValues[`${baseKey}-discordChannelId`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-discordChannelId`, text)}
                            />
                            <Text style={styles.inputLabel}>Message Content:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter message content"
                                value={inputValues[`${baseKey}-content`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-content`, text)}
                            />
                        </>
                    );
                case 'gmail':
                    return (
                        <>
                            <Text style={styles.inputLabel}>Recipient Email:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter recipient email"
                                value={inputValues[`${baseKey}-recipient`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-recipient`, text)}
                            />
                            <Text style={styles.inputLabel}>Email Subject:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter email subject"
                                value={inputValues[`${baseKey}-subject`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-subject`, text)}
                            />
                            <Text style={styles.inputLabel}>Email Body:</Text>
                            <TextInput
                                style={styles.input}
                                placeholder="Enter email body"
                                value={inputValues[`${baseKey}-body`] || ''}
                                onChangeText={(text) => handleInputChange(`${baseKey}-body`, text)}
                                multiline
                            />
                        </>
                    );
                default:
                    return <Text style={styles.noInputsText}>No specific inputs for this reaction provider.</Text>;
            }
        }
        return null;
    };

    return (
        <ScrollView style={styles.container}>
            {flow.map((flowItem, actionIndex) => (
                <View key={actionIndex} style={styles.flowItem}>
                    <TouchableOpacity onPress={() => handleSelectAction(actionIndex)}>
                        <View style={styles.actionContainer}>
                            <Text style={styles.actionText}>Provider: {flowItem.provider}</Text>
                            <Text style={styles.actionText}>Action: {flowItem.action.name}</Text>
                        </View>
                    </TouchableOpacity>

                    {flowItem.reactions.map((reaction, reactionIndex) => (
                        <TouchableOpacity key={reactionIndex} onPress={() => handleSelectReaction(actionIndex, reactionIndex)}>
                            <View style={styles.reactionContainer}>
                                <Text style={styles.reactionText}>Provider: {reaction.provider}</Text>
                                <Text style={styles.reactionText}>Reaction: {reaction.name}</Text>
                            </View>
                        </TouchableOpacity>
                    ))}

                    {selectedItem.actionIndex === actionIndex && selectedItem.type === 'action' && (
                        <View style={styles.infoCard}>
                            <Text style={styles.infoText}>Provider: {flowItem.provider}</Text>
                            <Text style={styles.infoText}>Action: {flowItem.action.name}</Text>
                            {renderInputs('action', flowItem.provider, actionIndex)}
                        </View>
                    )}

                    {selectedItem.actionIndex === actionIndex && selectedItem.type === 'reaction' && selectedItem.reactionIndex !== null && (
                        <View style={styles.infoCard}>
                            <Text style={styles.infoText}>Reaction Selected</Text>
                            <Text>{flowItem.reactions[selectedItem.reactionIndex!].provider}: {flowItem.reactions[selectedItem.reactionIndex!].name}</Text>
                            {renderInputs('reaction', flowItem.reactions[selectedItem.reactionIndex!].provider, selectedItem.reactionIndex!)}
                        </View>
                    )}

                    <View style={styles.buttonsContainer}>
                        <TouchableOpacity style={styles.actionButton} onPress={() => onAddReaction(actionIndex)}>
                            <MaterialIcons name="add" size={24} color="#fff" />
                            <Text style={styles.actionButtonTxt}>Add Reaction</Text>
                        </TouchableOpacity>
                        <View style={styles.options}>
                            <TouchableOpacity style={styles.actionButton} onPress={() => onRemoveAction(actionIndex)}>
                                <Entypo name="cross" size={24} color="#fff" />
                                <Text style={styles.actionButtonTxt}>Remove</Text>
                            </TouchableOpacity>
                            <TouchableOpacity style={styles.actionButton} onPress={() => handleSaveTrigger(actionIndex)}>
                                <MaterialIcons name="save-alt" size={24} color="#fff" />
                                <Text style={styles.actionButtonTxt}>Save</Text>
                            </TouchableOpacity>
                        </View>
                    </View>
                </View>
            ))}
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
    },
    flowItem: {
        marginBottom: 20,
        backgroundColor: Colors.light.tintLight,
        padding: 10,
        borderRadius: 10,
        elevation: 3,
    },
    actionContainer: {
        backgroundColor: Colors.light.tintDark,
        padding: 10,
        borderRadius: 5,
        marginBottom: 10,
    },
    actionText: {
        fontWeight: 'bold',
        color: '#fff',
    },
    reactionContainer: {
        backgroundColor: '#fff',
        padding: 10,
        borderRadius: 5,
        marginTop: 5,
    },
    reactionText: {
        fontWeight: 'bold',
        color: Colors.light.tintDark,
    },
    buttonsContainer: {
        flexDirection: 'column',
        marginTop: 10,
    },
    options: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    actionButton: {
        flexDirection: 'row',
        alignItems: 'center',
        backgroundColor: Colors.light.tint,
        padding: 10,
        borderRadius: 20,
        marginTop: 7.5,
    },
    actionButtonTxt: {
        color: '#fff',
        marginLeft: 5,
    },
    infoCard: {
        backgroundColor: '#f5f5f5',
        padding: 15,
        borderRadius: 8,
        marginTop: 30,
    },
    infoText: {
        fontWeight: 'bold',
        color: Colors.light.tint,
    },
    inputLabel: {
        fontWeight: 'bold',
        color: Colors.light.tintDark,
        marginTop: 10,
    },
    input: {
        backgroundColor: '#fff',
        borderRadius: 5,
        padding: 8,
        marginTop: 5,
        borderWidth: 1,
        borderColor: '#ddd',
    },
    noInputsText: {
        color: Colors.light.tintDark,
        fontStyle: 'italic',
        marginTop: 10,
    },
});
