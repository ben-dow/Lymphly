import {Box, Tabs, Text} from '@mantine/core'
import Radar from 'radar-sdk-js';
import 'radar-sdk-js/dist/radar.css'
import RadarMap from 'radar-sdk-js/dist/ui/RadarMap';
import { useEffect, useLayoutEffect, useState } from 'react';


export default function Search(){

    return(
        <div className="w-full h-full flex justify-center">
            <SearchContainer/>
        </div>
    )
}

interface Practice {
    PracticeId: String
    Name: String
    Lat: number
    Long: number
}


function SearchContainer(){

    const [practices, setPractices] = useState<Practice[]>([])

    useEffect(() => {
        fetch("/api/v1/providersearch/practices/all").then((r) =>r.json()).then(j=>{
            const pr: MapProps = j
            setPractices(pr.Practices)
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
                    <Map Practices={practices}/>
                </Tabs.Panel>
                <Tabs.Panel value="Providers" className='h-full'>
                    Providers
                </Tabs.Panel>
                <Tabs.Panel value="Practices" className='h-full'>
                    Practices
                </Tabs.Panel>
            </Tabs>
            </Box>

        </Box>            
    )
}



interface MapProps {
    Practices: Practice[]
}

function Map(props:MapProps){

    const [map, setMap] = useState<RadarMap>()

    useEffect(() => {

        fetch("/radar_pub_key.txt").then((r) =>r.text()).then(text=>{
            Radar.initialize(text);
            const Map = Radar.ui.map({
                container: "map",
                center: [-98.5556199, 39.8097343],
                zoom: 2,
            })
            setMap(Map)
        })

    }, [])

    useEffect(() => {
        if (map != undefined){
            for (let index = 0; index < props.Practices.length; index++) {
                const element = props.Practices[index];
                Radar.ui.marker({
                    color: '#000257',
                    width: 5,
                    height: 10,
                }).setLngLat([element.Long, element.Lat]).addTo(map)
            }
        }
      }, [props, map]);

      return(
        <Box className='w-full h-full'>
            <div id="map" className="w-full h-87 rounded-2xl"/>
        </Box>        
      )

}