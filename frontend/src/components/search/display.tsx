import { Box, Button, Divider, Table, TableData, Tabs } from "@mantine/core";
import Radar from "radar-sdk-js";
import RadarMap from "radar-sdk-js/dist/ui/RadarMap";
import { useCallback, useEffect, useState } from "react";
import { LimitedPracticePracticeListI as LimitedPracticesListI, PracticeI, PracticeListI as PracticeListI, ProviderListI } from "../../model/practice";
import { LngLatBoundsLike, LngLatLike} from "maplibre-gl";
import {Position} from "geojson"

export function DataDisplay() {
    const [selectedPractice, setSelectedPractice] = useState("")
    const [practices, setPractices] = useState<LimitedPracticesListI>({practices:[]})
    const [mapCfg, setMapCfg] = useState<MapConfiguration>()

    return (
        <Box className="flex flex-col w-full md:w-7xl">
            <Box className="w-full flex justify-center rounded-tl-xl bg-cyan-500">
                <Tabs defaultValue={"browse"} className="w-full" orientation='vertical' keepMounted={false}>
                    <Tabs.List className="bg-cyan-600 rounded-t-xl md:rounded-tl-xl md:rounded-tr-none">
                        <Tabs.Tab value="browse"><h1 className="font-sans text-white font-medium">Browse</h1></Tabs.Tab>
                        <Tabs.Tab value="current"><h1 className="font-sans text-white font-medium">Current Location</h1></Tabs.Tab>
                        <Tabs.Tab value="addr"><h1 className="font-sans text-white font-medium">Address</h1></Tabs.Tab>
                        <Tabs.Tab value="state"><h1 className="font-sans text-white font-medium">State</h1></Tabs.Tab>
                    </Tabs.List>
                    <Tabs.Panel value="browse" className='flex justify-center w-full h-full'>
                        <Browse setMapCfg={setMapCfg}  setPractices={setPractices}/>
                    </Tabs.Panel>
                    <Tabs.Panel value="current" className='flex justify-center w-full h-full'>
                        <SearchByLocation setMapCfg={setMapCfg} setPractices={setPractices}/>
                    </Tabs.Panel>
                    <Tabs.Panel value="addr" className='h-full w-full'>
                        Address
                    </Tabs.Panel>
                    <Tabs.Panel value="state" className='h-full w-full'>
                        State
                    </Tabs.Panel>
                </Tabs>
            </Box>
            <Box className="flex flex-col md:flex-row w-full justify-center">
                <Tabs defaultValue={"Map"} className="w-full md:w-7/8 max-w-5xl">
                    <Tabs.List className="bg-cyan-700  md:rounded-tr-none">
                        <Tabs.Tab value="Map"><h1 className="font-sans text-white font-medium">Map</h1></Tabs.Tab>
                        <Tabs.Tab value="Table"><h1 className="font-sans text-white font-medium">Table</h1></Tabs.Tab>
                    </Tabs.List>
                    <Tabs.Panel value="Map" className='flex justify-center w-full h-full'>
                        <Map mapConfiguration={mapCfg} setSelectedPractice={setSelectedPractice} practiceList={practices}/>
                    </Tabs.Panel>
                    <Tabs.Panel value="Table" className='h-full w-full'>
                        <PracticeTable practiceList={practices} updatedSelected={setSelectedPractice}/>
                    </Tabs.Panel>
                </Tabs>
                <Box className="md:h-full md:w-xs min-h-25 bg-cyan-950 ">
                    <Selected practiceId={selectedPractice}/>
                </Box>
            </Box>
            </Box>
    )
}

interface BrowseProps {
    setPractices: (props: LimitedPracticesListI) => void
    setMapCfg: (mapCfg: MapConfiguration) => void
}

function Browse(props: BrowseProps){

    useEffect(()=>{
        fetch("/api/v1/providersearch/practices/all").then((r) =>r.json()).then(j=>{
            const pl: PracticeListI = j
            props.setPractices(pl)
            props.setMapCfg({RadiusFeature:false})
        })
    }, [])


    return (
        <Box>
            Browse using the map or table below
        </Box>
    )
}

interface PracticeUpdaterI{
    setPractices: (props: LimitedPracticesListI) => void
    setMapCfg: (mapCfg: MapConfiguration) => void

}


function SearchByLocation(props:PracticeUpdaterI){

    const [radius, setRadius] = useState(25)

    useEffect(()=>{
        navigator.geolocation.getCurrentPosition((pos)=>{
            let lat = pos.coords.latitude
            let long = pos.coords.longitude

            fetch(`/api/v1/providersearch/practices/locate/proximity?lat=${lat}&long=${long}&radius=${radius}`).
                then(res => res.json()).
                then((res)=>{
                    props.setPractices(res)
                    props.setMapCfg({
                        RadiusFeature: true,
                        RadiusOrigin: [long, lat],
                        Radius: radius
                    })
                })
            
        })
    }, [radius])


    return (
        <div>
           <Box className='flex justify-center h-full flex-col p-2'>
                <Box className="bg-white p-4 flex flex-col gap-2">
                    <Box className="flex flex-col sm:flex-row gap-5 ">
                        <Box className='font-sans text-xl font-medium text-sky-950 flex flex-col justify-center'>Search Radius:</Box>
                        <Button onClick={()=>{setRadius(25)}}>25 Miles  </Button>
                        <Button onClick={()=>{setRadius(50)}}>50 Miles  </Button>
                        <Button onClick={()=>{setRadius(100)}}>100 Miles  </Button>
                    </Box>
                    <Box className="text-center">
                        Current: {radius} Miles
                    </Box>
                </Box>
            </Box>
        </div>
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


interface PracticeTableProps {
    practiceList: LimitedPracticesListI
    updatedSelected: (practiceId :string ) => void
}

export function PracticeTable(props: PracticeTableProps){
    let rows: JSX.Element[] = []

    rows = props.practiceList.practices.map((r, idx)=>{
        return (
            <Table.Tr key={idx}>
                <Table.Td onClick={ ()=>{props.updatedSelected(r.practiceId)}} className=" hover:bg-cyan-600 hover:cursor-pointer font-sans font-medium text-white">{r.name}</Table.Td> 
            </Table.Tr>
        )
    })

    return (
        <Box className="w-full h-75 md:h-150 md:rounded-bl-2xl bg-cyan-900 p-5 overflow-auto">
            <Table withTableBorder>
                <Table.Thead >
                    <Table.Tr>
                        <Table.Th className="font-sans font-bold text-lg text-white">Name</Table.Th>
                    </Table.Tr>
                </Table.Thead>
                <Table.Tbody >
                    {rows}
                </Table.Tbody>
            </Table>
    </Box>
    )

}

interface MapProps {
    mapConfiguration?:MapConfiguration
    practiceList: LimitedPracticesListI
    setSelectedPractice: (practiceId: string) => void
}

interface MapConfiguration {
    RadiusOrigin?: LngLatLike
    RadiusFeature?: boolean
    Radius?: number
}

export function Map(props: MapProps){
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
                map.addPolygon({
                    type: "Feature",
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