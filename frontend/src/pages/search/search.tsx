import {Box, Divider, Tabs, Text} from '@mantine/core'
import 'radar-sdk-js/dist/radar.css'
import { useEffect, useState } from 'react';
import { PracticeListI } from '../../model/practice';
import { DataDisplay, Map } from '../../components/search/display';
import { Routes, Route } from 'react-router';


export default function Search(){

    return(
        <div className="w-full h-full flex justify-center">
            <Box className={"bg-sky-200 p-5 shadow-sm w-full h-full"}>
                <Routes>
                    <Route path="/" element={<SearchHome/>}/>
                </Routes>
            </Box>
        </div>
    )
}

function SearchHome(){
    const [practices, setPractices] = useState<PracticeListI>({practices:[]})

    useEffect(() => {
        fetch("/api/v1/providersearch/practices/all").then((r) =>r.json()).then(j=>{
            const pl: PracticeListI = j
            setPractices(pl)
        })
    }, [])

    return(
        <Box className="flex flex-col gap-5">
            <Box className="text-center">
                <h2 className='text-3xl font-normal font-sans'>Find a Provider</h2>
            </Box>
            <Box className='flex flex-row flex-wrap justify-center gap-5'>
                <Box className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>By Current Location</Text>
                </Box>

                <Box className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>By Address</Text>
                </Box>

                <Box className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>By State</Text>
                </Box>
            </Box>

            <Box className='flex flex-row flex-wrap justify-center gap-5'>
                <Divider className={"w-full max-w-2xl"} color={"black"}/>
            </Box>
            <Box className='flex flex-row flex-wrap justify-center gap-5'>
                <Box className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>Browse</Text>
                </Box>
            </Box>

            <Box className='flex justify-center'>
                <Map practiceList={practices}/>
            </Box>

        </Box>            
    )
}




