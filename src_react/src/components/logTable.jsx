import React from 'react'
import {
    Table,
    Thead,
    Tbody,
    Tfoot,
    Tr,
    Th,
    Td,
    TableCaption,
    Box
  } from "@chakra-ui/react"

/*
    This component displays the log in formatted display
*/
class LogTable extends React.Component {
    render() {
        return (<Box maxW='50%'>
                <Table size='sm'>
                <Thead>
                    <Tr>
                        <Td>Time</Td>
                        <Td>Level</Td>
                        <Td>Package</Td>
                        <Td>Message</Td>
                        <Td>Category</Td>
                    </Tr>
                </Thead>
                <Tbody>
                    <Tr>
                        <Td>Oct 20</Td>
                        <Td>Debug</Td>
                        <Td>ela.system</Td>
                        <Td>This is sample</Td>
                        <Td>none</Td>
                    </Tr>
                </Tbody>
                <Tfoot>

                </Tfoot>
            </Table>
        </Box>)
    }
}

export default LogTable