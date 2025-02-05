import {Box, Tabs, Text} from '@mantine/core'
import Radar from 'radar-sdk-js';
import 'radar-sdk-js/dist/radar.css'
import { useEffect } from 'react';


export default function Search(){

    return(
        <div className="w-full h-full flex justify-center">
            <SearchContainer/>
        </div>
    )
}


function SearchContainer(){



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
            
            <Box className="text-center">
                <Text size="xl" fw={700}>Or View All Providers</Text>
            </Box>

            <Box>                
            <Tabs defaultValue={"Map"} className="w-full h-96 shadow-sm  rounded-2xl">
                <Tabs.List>
                    <Tabs.Tab value="Map">Map</Tabs.Tab>
                    <Tabs.Tab value="Providers">Providers</Tabs.Tab>
                    <Tabs.Tab value="Practices">Practices</Tabs.Tab>
                </Tabs.List>
                <Tabs.Panel value="Map" className='h-full'>
                    <Map/>
                </Tabs.Panel>
                <Tabs.Panel value="Providers" className='h-full'>
                    <Map/>
                </Tabs.Panel>
                <Tabs.Panel value="Practices" className='h-full'>
                    <Map/>
                </Tabs.Panel>
            </Tabs>
            </Box>

        </Box>            
    )
}


function Map(){
    useEffect(() => {
        fetch("/radar_pub_key.txt").then((r) =>r.text()).then(text=>{
            Radar.initialize(text);
            Radar.ui.map({
                container: "map",
                center: [-98.5556199, 39.8097343],
                zoom: 2,
            })
        })

      }, []);

      return(
        <Box className='w-full h-full'>
            <div id="map" className="w-full h-87 rounded-2xl"/>
        </Box>        
      )

}