import React from 'react'
import { Box, Flex } from '@chakra-ui/react'
import LogView from './logTable'
import Filter from './filterview'
import {EboxEvent} from 'elabox-foundation'
import  { CODE_SUCCESS } from '../constant'
import { retrieveLatest, retrieveSummary } from '../actions'

class LogArea extends React.Component {
    state = { loading: true, summary: {}, logs: [], filter: null }
    componentDidMount() {
        this.initLogs()
    }
    async initLogs() {
        const { eventH = new EboxEvent('http://' + window.location.hostname)} = this.props
        this.setState({loading: true, eventh:eventH})
        eventH.waitUntilConnected() 
            .then(connected => {
                retrieveSummary(eventH, summary => {
                    console.log(summary)
                    if (summary.code === CODE_SUCCESS)
                        this.setState({summary:  summary.message, loading: false})
                    else 
                        this.onError(summary.message)
                })
                retrieveLatest(eventH, this.state.filter, logs => {
                    this.setState({logs: logs.message})
                })
            }).catch(err => {
                this.onError(err) 
                this.setState({loading: false})
            })
    }
    onError(err) {
        console.log("ERRR ", err)
    }
    onChangedFilter(filter) {
        console.log(filter)
        this.setState({filter: filter})
        retrieveLatest(this.state.eventh, filter, logs => {
            console.log(logs)
            this.setState({logs:logs.message}) 
        })
    }
    render() {
        const { loading, summary, logs } = this.state
        if (loading) {
            return "Loading"
        }
        return (
            <Flex h='700px' w='1200px' >
                <Box minW='300px'h='100%' bg='gray.300'>
                    <Filter summary={summary} onChanged={this.onChangedFilter.bind(this)}/>
                </Box>
                <Flex flex='1' flexFlow='column' w='900px'>
                    {/* <Box h='calc(20vh)' ></Box> */}
                    <Box flex='1' ><LogView logs={logs}/></Box>
                </Flex>
            </Flex>
        )
    }
}
export default LogArea