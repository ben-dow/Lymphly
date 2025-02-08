import {Box, Tabs, Text} from '@mantine/core'
import 'radar-sdk-js/dist/radar.css'
import { useEffect, useState } from 'react';
import { PracticeList } from '../model/practice';
import { DataDisplay, Map } from '../components/search/display';


export default function Search(){

    return(
        <div className="w-full h-full flex justify-center">
            <SearchContainer/>
        </div>
    )
}




function SearchContainer(){

    const [practices, setPractices] = useState<PracticeList>({practices:[]})

    useEffect(() => {
        fetch("/api/v1/providersearch/practices/all").then((r) =>r.json()).then(j=>{
            const pl: PracticeList = j
            setPractices(pl)
        })
    }, [])

    return(
        <Box className={"bg-emerald-100 p-5 flex flex-col gap-10 shadow-sm w-full xl:w-1/2 p-10 rounded-xl"}>
            <Box className="text-center">
                <Text size="xl" fw={700}>Find A Provider By:</Text>
            </Box>

            <Box className='flex flex-row flex-wrap justify-center gap-5'>
                <Box className='rounded-xl h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>Current Location</Text>
                </Box>

                <Box className='rounded-xl h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>State</Text>
                </Box>

                <Box className='rounded-xl h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>City</Text>
                </Box>

                <Box className='rounded-xl h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>Practice</Text>
                </Box>
            </Box>

            <Box>
                <DataDisplay practiceList={practices}/>
            </Box>

        </Box>            
    )
}




