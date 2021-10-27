import React from 'react'
import { Box, Flex } from '@chakra-ui/react'
import LogView from './logTable'
import Filter from './filterview'
import {EboxEvent} from 'elabox-foundation'
import  { CODE_SUCCESS } from '../constant'
import { retrieveLatest, retrieveSummary } from '../actions'

class LogArea extends React.Component {
    state = { 
        loading: true, 
        offset: 0, // the last position from log file retrieved. 0 if start of most recent log
        summary: {}, 
        logs: [], 
        filter: null, 
        loadingLogs: false, 
        loadingPrevious: false }
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
                this.onRetrieveLogs(0)
            }).catch(err => {
                this.onError(err) 
                this.setState({loading: false})
            })
    }
    // use to retrieve log, 
    // @startoffset the offset from which the log will start to retrieve. starting from end of file
    // @appendLogs - true if will concatenate the log to end
    onRetrieveLogs(startOffset, appendLogs = false) {
        if (this.state.loadingLogs) return;
        this.setState({loadingLogs: true})
        retrieveLatest(this.state.eventh, {...this.state.filter, offset:startOffset}, logs => {
            //console.log(logs)
            if (logs.code === CODE_SUCCESS){
                var newLogs = logs.message.logs
                if (appendLogs) {
                    newLogs = [...this.state.logs, ...newLogs]
                }
                //console.log(newLogs)
                this.setState({logs: newLogs, offset: this.state.offset + logs.message.size, loadingLogs: false})
            }else {
                this.onError(logs.message)
                this.setState({loadingLogs: false})
            }
        })
    }
    onRetrievePrevious() {
        //console.log("retrieve previous " + this.state.offset)
        this.onRetrieveLogs(this.state.offset, true)
    }
    onRetrieveLatest() {
        //console.log("retrieve latest")
        this.onRetrieveLogs(0)
    }
    onError(err) {
        console.log("ERRR ", err)
    }
    onChangedFilter(filter) {
        console.log(filter)
        this.setState({filter: filter})
        retrieveLatest(this.state.eventh, filter, logs => {
            console.log(logs)
            this.setState({logs:logs.message.logs, offset: 0}) 
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
                    <Box flex='1' ><LogView logs={logs} 
                        onLatest={this.onRetrieveLatest.bind(this)} 
                        onPrevious={this.onRetrievePrevious.bind(this)}/></Box>
                </Flex>
            </Flex>
        )
    }
}
export default LogArea