import {Box, Tabs, Text} from '@mantine/core'
import Radar from 'radar-sdk-js';
import 'radar-sdk-js/dist/radar.css'
import { useEffect } from 'react';


export default function Search(){
    Radar.initialize('prj_test_pk_16c27944fa53ebb83144216b8649c0e6e2f9db92');

    return(
        <div className="w-full h-full flex justify-center">
            <SearchContainer/>
        </div>
    )
}


function SearchContainer(){

    useEffect(() => {
        Radar.ui.map({
            container: "map",
            center: [-98.5556199, 39.8097343],
            zoom: 2,
        })
      }, []);

    return(
        <Box className={"bg-emerald-100 p-5 flex flex-col gap-10 shadow-sm max-w-7xl p-10 rounded-xl"}>
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
                
            <Tabs defaultValue={"Map"} className="w-full h-96 shadow-sm  rounded-2xl">
                <Tabs.List>
                    <Tabs.Tab value="Map">Map</Tabs.Tab>
                </Tabs.List>
                <Tabs.Panel value="Map" className='h-full'>
                    <Box className='w-full h-full'>
                        <div id="map" className="w-full h-87 rounded-2xl"/>
                    </Box>                
                </Tabs.Panel>
            </Tabs>
            </Box>

        </Box>            
    )
}
