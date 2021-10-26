
import { Box,
    Accordion,
    AccordionButton,
    AccordionPanel,
    AccordionItem,
    AccordionIcon,
    Button,
    HStack,
    VStack,
    Text,
    Input,
    Select,
    Container,
    Wrap,
    Tag,
    TagLabel,
    TagCloseButton
} from '@chakra-ui/react'
import { CloseIcon } from '@chakra-ui/icons'
import { useState } from 'react'

const emptyFilter = {
    levels: {},
    packages: {},
    categories: {},
    conditions: []
}

const ItemButton = ({
    children, 
    toggle = true, 
    onToggle = (val) => {}, 
    color = "gray",
    counter = -1
}) => {
    const [currentToggle, setToggle] = useState(toggle)
    const updateToggle = (newVal) => {
        onToggle(newVal)
        setToggle(newVal)
    }
    
    return (
        <Tag size="md"
        onClick={() => updateToggle(!currentToggle)}
        borderRadius="full"
        variant={currentToggle ? 'solid' : 'outline'}
        colorScheme={color}>
            <HStack>
            <TagLabel>{children}</TagLabel>
            {counter > 1 && <Text fontSize="sm" variant='outline'> {"(" + counter + ")"}</Text>}
            </HStack>
            {currentToggle &&<TagCloseButton/>}
        </Tag>
    )
}

const FilterGroup = ({label,children}) => {
    return (
        <AccordionItem>
            <AccordionButton>
                <Box flex='1' textAlign='left'>{label}</Box>
                <AccordionIcon/>
            </AccordionButton>
            <AccordionPanel>
                <Wrap spacing='10px'>
                    {children.length > 0 && children}
                    {!(children.length > 0) && <Tag variant='outline'>Empty</Tag>}
                </Wrap>
            </AccordionPanel>
        </AccordionItem>
    )
}

// display filter for conditions
const ConditionFilter = ({conditions = [], onChanged=(newConditions)=>{}}) => {
    const emptyCon = {key:'', value: '',  operator: '', on:true}
    const [currentCons, setConditions] = useState(conditions)
    const [currentCon, setCondition] = useState({...emptyCon})
    const addCondition = () => {
        if (!currentCon.key || currentCon.key === '' ||
            !currentCon.value || currentCon.value === '' ||
            !currentCon.operator || currentCon.operator === '')
            return
        currentCons.push(currentCon)
        onChanged(currentCons)
        setConditions(currentCons)
        setCondition({...emptyCon})
    }
    const removeCondition = (index) => {
        const newConst = [...currentCons]
        newConst.splice(index, 1)
        onChanged(newConst)
        setConditions(newConst)
    }
    return (
        <FilterGroup label="Conditions">
            <Container><Text fontSize='sm'>Need to satisfy all conditions.</Text></Container>
            {
                currentCons.map((val, index) => (
                    <ItemButton key={val.key+val.operator+index} toggle={val.on}
                        onToggle={(toggle) => {
                            removeCondition(index)
                        }}>
                        {val.key + ' ' + val.operator + ' ' + val.value}
                    </ItemButton> 
                ))
            }

            <VStack size='sm'>
                <Input size='sm' placeholder='field' value={currentCon.key} onChange={(event) => {
                    setCondition({...currentCon, key:event.target.value})
                }}/>
                <Select size='sm' placeholder='select condition' value={currentCon.operator} onChange={(event) => {
                    setCondition({...currentCon, operator:event.target.value})
                }}>
                    <option value="contains">contains</option>
                    <option value="not contains">not contains</option>
                    <option value="==">==</option>
                    <option value="!=">!=</option>
                </Select>
                <Input size='sm' placeholder='value' value={currentCon.value} onChange={(event) => {
                    setCondition({...currentCon, value:event.target.value})
                }}/>
                <Button variant='outline' size='sm' onClick={()=> addCondition()}>
                    Add Condition
                </Button>
            </VStack>
        </FilterGroup>
    )
}
 
// main component for rendering filter view
// @summary the log summary prop
const filter = ({summary = {}, filter = emptyFilter, onChanged = (newVal) => {}}) => {
    const newData = {...filter}
    //newData.conditions = [{key:'message', operator: 'contains', value:'hello', on: true}]
    const onDataChanged = () => {
        onChanged(newData)
    }
    return (
        <Box>
            <Accordion allowToggle allowMultiple>
                <FilterGroup label="Level">
                    <ItemButton color='green' counter={summary.levels.info} onToggle={(toggle) => {
                            newData.levels["info"] = toggle; 
                            onDataChanged()    
                        }}>
                        Info
                    </ItemButton>
                    <ItemButton counter={summary.levels.debug} onToggle={(toggle) => {
                            newData.levels["debug"] = toggle; 
                            onDataChanged()    
                        }}>
                        Debug
                    </ItemButton>
                    <ItemButton color='yellow' counter={summary.levels.warning}onToggle={(toggle) => {
                            newData.levels["warning"] = toggle; 
                            onDataChanged()    
                        }}>
                        Warning
                    </ItemButton>
                    <ItemButton color='red' counter={summary.levels.error} onToggle={(toggle) => {
                            newData.levels["error"] = toggle; 
                            onDataChanged()    
                        }}>
                        Error
                    </ItemButton>
                </FilterGroup>
                <FilterGroup label="Package">
                    {
                        Object.keys( summary.packages).map((key, index) => (
                            <ItemButton 
                                counter={summary.packages[key]}
                                key={key + index}
                                toggle={newData.packages[key]} 
                                onToggle={(toggle) => {
                                    newData.packages[key] = toggle; 
                                    onDataChanged()    
                                }}>
                                {key}
                            </ItemButton>
                        ))
                    }
                </FilterGroup>
                <FilterGroup label="Category">
                    {
                        Object.keys(summary.categories).map((key, index) => (
                            <ItemButton 
                                counter={summary.categories[key]}
                                key={key + index}
                                toggle={newData.categories[key]} 
                                onToggle={(toggle) => {
                                    newData.categories[key] = toggle; 
                                    onDataChanged()    
                                }}>
                                {key}
                            </ItemButton>
                        ))
                    }
                </FilterGroup>
                <ConditionFilter conditions={newData.conditions} onChanged={
                    (newVal) => {
                        newData.conditions = newVal
                        onDataChanged()
                    }
                }/>
            </Accordion>
        </Box>
    )
}

export default filter