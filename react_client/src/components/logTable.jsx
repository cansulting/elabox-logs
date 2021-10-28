import React from 'react'
import {
    Box,
    Container,
    Accordion,
    AccordionItem,
    AccordionPanel,
    AccordionButton,
    AccordionIcon,
    HStack,
  } from "@chakra-ui/react"

/*
    This component displays the log in formatted display
*/

const CHARLIMIT = 50

function getColor(level) {
    switch (level) {
        case 'error':
            return 'red'
        case 'warning': 
            return 'violet'
        case 'debug':
            return 'blue.500'
        default:
            return "";
    }
}

class LogTable extends React.Component {
    lastScroll=0
    getScrollPercent() {
        const logsElem = document.getElementById("logs")
        const h = logsElem.clientHeight
        const scrollH = logsElem.scrollHeight - h
        const percent = Math.abs( logsElem.scrollTop / scrollH)
        return percent
    }
    onScroll() {
        const { onLatest, onPrevious } = this.props
        const percent = this.getScrollPercent()
        //console.log(this.lastScroll - percent,percent)
        if (percent <= 0.01) {
            // check direction
            if (this.lastScroll - percent > 0)
                onLatest()
        } else if (percent >= 0.99) {
            if (this.lastScroll - percent < 0)
                onPrevious()
        }
        this.lastScroll = percent
        
    }
    render() {
        const  { logs = [] } = this.props
        
        return (
            <Box flex='1'>
                <HStack fontWeight='semibold'>
                    <Container w='container.xs'>Level</Container>
                    <Container w='container.sm'>Time</Container>
                    <Container w='container.sm'>Package</Container>
                    <Container w='container.xl'>Message</Container>
                    <Container w='container.sm'>Category</Container>
                </HStack>
                <Accordion id="logs" allowToggle overflowY='auto' h='calc(100vh - 240px)' display='flex' flexDirection='column-reverse' onScroll={this.onScroll.bind(this)}>
                    {
                        logs.map( (val, index) => (
                            val.level && <AccordionItem key={"logv" + index}>
                                <AccordionButton textColor={getColor(val.level)}>
                                    <AccordionIcon/>
                                    <Container w='container.xs'>{val.level}</Container>
                                    <Container w='container.sm'>{val.time}</Container>
                                    <Container w='container.sm'>{val.package}</Container>
                                    <Container w='6xl'>{val.message.substr(0, CHARLIMIT)+'...'}</Container>
                                    <Container w='container.sm'>{val.category}</Container>
                                </AccordionButton>
                                <AccordionPanel pb='4' textAlign='left'>
                                    <h2><pre>{JSON.stringify(val, null, '\t')}</pre></h2>
                                </AccordionPanel>
                            </AccordionItem>
                        ))
                    }
                </Accordion>
        </Box>)
    }
}

export default LogTable