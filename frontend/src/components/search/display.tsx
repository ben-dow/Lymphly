import { Box, Divider, Table, TableData, Tabs } from "@mantine/core";
import Radar from "radar-sdk-js";
import RadarMap from "radar-sdk-js/dist/ui/RadarMap";
import { useCallback, useEffect, useState } from "react";
import { LimitedPracticePracticeListI as LimitedPracticesListI, PracticeI, PracticeListI as PracticeListI, ProviderListI } from "../../model/practice";
import { LngLatBoundsLike, LngLatLike} from "maplibre-gl";
import {Position} from "geojson"


export interface DataDisplayProps {
    practiceList?: PracticeListI
    mapConfiguration?: MapConfiguration
    setSelectedPractice?: (id:string)=>void
}

export function DataDisplay(props:DataDisplayProps) {
    const [selectedPractice, setSelectedPractice] = useState("")

    return (
            <Box className="flex flex-col md:flex-row w-full justify-center">
                <Tabs defaultValue={"Map"} className="w-full md:w-7/8 max-w-5xl">
                    <Tabs.List className="bg-cyan-700 rounded-t-xl md:rounded-tl-xl md:rounded-tr-none">
                        <Tabs.Tab value="Map"><h1 className="font-sans text-white font-medium">Map</h1></Tabs.Tab>
                        <Tabs.Tab value="Map"><h1 className="font-sans text-white font-medium">Practices</h1></Tabs.Tab>
                    </Tabs.List>
                    <Tabs.Panel value="Map" className='flex justify-center w-full h-full'>
                        <Map setSelectedPractice={setSelectedPractice} {...props}/>
                    </Tabs.Panel>
                    <Tabs.Panel value="Practices" className='h-full w-full'>
                        <PracticeTable {...props}/>
                    </Tabs.Panel>
                </Tabs>
                <Box className="md:h-full md:w-xs min-h-25 bg-cyan-700 ">
                    <Selected practiceId={selectedPractice}/>
                </Box>
            </Box>
    )
}

interface SelectedProps {
    practiceId: string
}

function Selected(props:SelectedProps){

    const [practice, setPractice] = useState<PracticeI>()
    const [providers, setProviders] = useState<ProviderListI>()

    useEffect(()=>{
        if (props.practiceId != ""){
            fetch("/api/v1/providersearch/practice/"+props.practiceId).then((r) =>r.json()).then(pr=>{
                setPractice(pr)
            })

            fetch("/api/v1/providersearch/practice/"+props.practiceId+"/providers").then((r) =>r.json()).then(pr=>{
                setProviders(pr)
            })
        }

    }, [props.practiceId])
    
    if (practice === undefined) {
        return (
            <Box className="text-center p-5 text-2xl text-white font-sans font-medium">No Practice Selected</Box>
        )
    } else {
        let tableData: TableData = {
            head: ["Name", "Tags"],
            body: []
        }

        if (providers != undefined){
            for(let i = 0; i<providers.providers.length; i++){
                tableData.body.push([providers.providers[i].name, providers.providers[i].tags])
            }    
        }

        

        return (
            <Box className="p-5 flex flex-col gap-5">
                <Box className="text-center font-sans text-2xl text-white font-medium">
                        Practice Info
                </Box>
                <Box className="text-white border-white border overflow-hidden font-sans flex flex-col gap-5 p-6 w-full justify-center text-wrap">
                   
                    <Box className="flex flex-row gap-2 justify-baseline flex-wrap">
                        <Box className="text-sm w-20 font-medium">Name: </Box>
                        <Box className="text-sm">{practice.name}</Box>
                    </Box>
                    <Box className="flex flex-row justify-baseline gap-2 flex-wrap">
                        <Box className="text-sm font-medium w-20">Address: </Box>
                        <Box className="text-sm">{practice.fullAddress}</Box>
                    </Box>
                    <Box className="flex flex-row justify-baseline gap-2 flex-wrap">
                        <Box className="text-sm font-medium w-20">Website: </Box>
                        <Box className="text-sm"><a  target="_blank" rel="noopener noreferrer" className="underline" href={practice.website}>{practice.website}</a></Box>
                    </Box>
                    <Box className="flex flex-row justify-baseline gap-2 flex-wrap">
                        <Box className="text-sm font-medium w-20 ">Phone: </Box>
                        <Box className="text-sm">{practice.phone}</Box>
                    </Box>
                    <Box className="flex flex-row justify-baseline gap-2 flex-wrap">
                        <Box className="text-sm font-medium w-20">Tags: </Box>
                        <Box className="text-sm">{practice.tags}</Box>
                    </Box>
                
                </Box>
                <Box className="text-center font-sans text-2xl text-white font-medium">
                        Providers
                </Box>
                <Table className="font-sans text-white" withTableBorder data={tableData}></Table>

            </Box>
        )
    }

  
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


interface MapConfiguration {
    RadiusOrigin?: LngLatLike
    RadiusFeature?: boolean
    Radius?: number
}

export function Map(props: DataDisplayProps){
    const [map, setMap] = useState<RadarMap>(undefined)

    useEffect(() => {
        if (map === undefined){
            fetch("/radar_pub_key.txt").then((r) =>r.text()).then(text=>{
                Radar.initialize(text, {debug: false});
                const Map = Radar.ui.map({
                    container: "map",
                    zoom: 0,
                })
                setMap(Map)
            })
        }

    }, [map])

    useEffect(() => {
        if (map != undefined){

            if (props.practiceList != undefined){
                map.clearMarkers()
                for (let index = 0; index < props.practiceList.practices.length; index++) {
                    const element = props.practiceList.practices[index];
                    const marker = Radar.ui.marker({
                        color: 'red',
                        scale: .5,
                        popup: {
                            text: element.name
                        },
                    })
                    marker.on("click", ()=>{if(props.setSelectedPractice !=undefined){props.setSelectedPractice(element.practiceId)}})
                    marker.setLngLat([element.longitude, element.lattitude]).addTo(map)
                }
                map.fitToMarkers()
            }
 
            map.clearFeatures()

            if (props.mapConfiguration != undefined && props.mapConfiguration.RadiusFeature) {
                const marker = Radar.ui.marker(
                    {
                        color: "blue",
                        scale: .75,
                    }
                )
                marker.setLngLat(props.mapConfiguration.RadiusOrigin).addTo(map)

                console.log(ZoneCoords(props.mapConfiguration.RadiusOrigin, props.mapConfiguration.Radius, 100))
                map.addPolygon({
                    type: "Feature",
                    id: 1,
                    properties: {
                        name: "radius"
                    },
                    geometry: {
                        type: "Polygon",
                        coordinates:ZoneCoords(props.mapConfiguration.RadiusOrigin, props.mapConfiguration.Radius, 100)
                    }
                }, {
                    paint: {
                        "fill-color": "yellow", 
                        "fill-opacity": .1,
                        "border-width": 1,
                        "border-color": "red",
                        "border-opacity": .3,
                    }
                })
                map.fitToFeatures()
            }

            map.redraw()
        }
      }, [props.practiceList, props.mapConfiguration,  map]);

      return(
        <div id="map" className="w-full h-75 md:h-150 md:rounded-b-2xl"/>
      )
}


function ZoneCoords(lngLt: LngLatLike, radius: number, resolution: number): Position[][] {
    const long = lngLt[0]
    const lat = lngLt[1]

    const radiusKm = radius / 0.621371;
    const radiusLon = 1 / (111.319 * Math.cos(lat * (Math.PI / 180))) * radiusKm;
    const radiusLat = 1 / 110.574 * radiusKm;
    
    const dTheta = 2 * Math.PI / resolution;
    let theta = 0;

    let out: Position[] = []

    for (var i = 0; i < resolution; i++)
    {
        out.push([long + radiusLon * Math.cos(theta),lat + radiusLat * Math.sin(theta)]);
        theta += dTheta;
    }


    return [out]
}