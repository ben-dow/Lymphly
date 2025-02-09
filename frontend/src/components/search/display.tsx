import { Box, Table, Tabs } from "@mantine/core";
import Radar from "radar-sdk-js";
import RadarMap from "radar-sdk-js/dist/ui/RadarMap";
import { useEffect, useState } from "react";
import { LimitedPracticePracticeListI as LimitedPracticesListI, PracticeListI as PracticeListI } from "../../model/practice";


interface DataDisplayProps {
    practiceList: PracticeListI
}

export function DataDisplay(props:DataDisplayProps) {
    return (
            <Tabs defaultValue={"Map"} className="w-full h-96 shadow-sm  rounded-2xl">
                <Tabs.List>
                    <Tabs.Tab value="Map">Map</Tabs.Tab>
                    <Tabs.Tab value="Practices">Practices</Tabs.Tab>
                </Tabs.List>
                <Tabs.Panel value="Map" className='h-full'>
                    <Map {...props}/>
                </Tabs.Panel>
                <Tabs.Panel value="Practices" className='h-full'>
                    <PracticeTable {...props}/>
                </Tabs.Panel>
            </Tabs>
    )
}


export function PracticeTable(props: DataDisplayProps){
    console.log("test")

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

export function Map(props: MapProps){
    const [map, setMap] = useState<RadarMap>()
    const practiceList = props.practiceList.practices

    useEffect(() => {
        fetch("/radar_pub_key.txt").then((r) =>r.text()).then(text=>{
            Radar.initialize(text);
            const Map = Radar.ui.map({
                container: "map",
                center: [-98.5556199, 39.8097343],
                zoom: 1,
            })
            setMap(Map)
        })

    }, [])

    useEffect(() => {
        map.clearMarkers()
        if (map != undefined){
            for (let index = 0; index < practiceList.length; index++) {
                const element = practiceList[index];
                Radar.ui.marker({
                    color: '#000257',
                    scale: .5,
                    popup: {
                        text: element.name
                    }
                }).setLngLat([element.longitude, element.lattitude]).addTo(map)
            }
        }
      }, [practiceList, map]);

      return(
        <Box className='w-full h-full'>
            <div id="map" className="w-full h-87 rounded-2xl"/>
        </Box>        
      )
}