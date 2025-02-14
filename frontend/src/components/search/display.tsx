import { Box, Table, Tabs } from "@mantine/core";
import Radar from "radar-sdk-js";
import RadarMap from "radar-sdk-js/dist/ui/RadarMap";
import { useCallback, useEffect, useState } from "react";
import { LimitedPracticePracticeListI as LimitedPracticesListI, PracticeListI as PracticeListI } from "../../model/practice";


interface DataDisplayProps {
    practiceList: PracticeListI
}

export function DataDisplay(props:DataDisplayProps) {
    return (
            <Tabs defaultValue={"Map"} className="rounded-2xl w-6xl h-150">
                <Tabs.List className="bg-cyan-700 rounded-t-xl">
                    <Tabs.Tab value="Map"><h1 className="font-sans text-white font-medium">Map</h1></Tabs.Tab>
                    <Tabs.Tab value="Map"><h1 className="font-sans text-white font-medium">Practices</h1></Tabs.Tab>
                </Tabs.List>
                <Tabs.Panel value="Map" className='flex justify-center w-full h-full'>
                    <Map {...props}/>
                </Tabs.Panel>
                <Tabs.Panel value="Practices" className='h-full w-full'>
                    <PracticeTable {...props}/>
                </Tabs.Panel>
            </Tabs>
    )
}


export function PracticeTable(props: DataDisplayProps){
    return (
    <Table>
        <Table.Thead>
            <Table.Tr>
                <Table.Th>Name</Table.Th>
                <Table.Th>Name</Table.Th>
            </Table.Tr>
        </Table.Thead>
        <Table.Tbody>
            <Table.Tr>
                <Table.Td>T</Table.Td>
            </Table.Tr>
        </Table.Tbody>
    </Table>)

}

interface MapProps{
    practiceList: LimitedPracticesListI
}

function useRadar() {
    const [initialized, setInitialized] = useState(false)

    const initialize = useCallback(()=>{
        fetch("/radar_pub_key.txt").then((r) =>r.text()).then(text=>{
            Radar.initialize(text);
            setInitialized(true);
        })
    }, [])

    useEffect(()=>{
        initialize()    
    },[])
    
    return initialized
}

export function Map(props: MapProps){
    let initialized = useRadar()
    const [map, setMap] = useState<RadarMap>()
    const practiceList = props.practiceList.practices

    useEffect(() => {
        if (initialized){
            const Map = Radar.ui.map({
                container: "map",
                zoom: 1,
                
            })
            setMap(Map)
        }
    }, [initialized])

    useEffect(() => {
        if (map != undefined){
            map.clearMarkers()
            for (let index = 0; index < practiceList.length; index++) {
                const element = practiceList[index];
                Radar.ui.marker({
                    color: 'red',
                    scale: .5,
                    popup: {
                        text: element.name
                    },
                    zoom: 1,
                }).setLngLat([element.longitude, element.lattitude]).addTo(map)
            }
            map.fitToMarkers()
        }
      }, [practiceList, map]);

      return(
        <div id="map" className="w-full h-full rounded-b-2xl"/>
      )
}