import { Box, Table, Tabs } from "@mantine/core";
import Radar from "radar-sdk-js";
import RadarMap from "radar-sdk-js/dist/ui/RadarMap";
import { useCallback, useEffect, useState } from "react";
import { LimitedPracticePracticeListI as LimitedPracticesListI, PracticeListI as PracticeListI } from "../../model/practice";
import { LngLatBoundsLike, LngLatLike} from "maplibre-gl";
import {Position} from "geojson"


export interface DataDisplayProps {
    practiceList?: PracticeListI
    mapConfiguration?: MapConfiguration
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


interface MapConfiguration {
    RadiusOrigin?: LngLatLike
    RadiusFeature?: boolean
    Radius?: number
}

export function Map(props: DataDisplayProps){
    const [map, setMap] = useState<RadarMap>(undefined)

    useEffect(() => {
            fetch("/radar_pub_key.txt").then((r) =>r.text()).then(text=>{
                Radar.initialize(text);
                const Map = Radar.ui.map({
                    container: "map",
                    zoom:0,
                })
                setMap(Map)
            })
    }, [])

    useEffect(() => {
        if (map != undefined){

            if (props.practiceList != undefined){
                map.clearMarkers()
                for (let index = 0; index < props.practiceList.practices.length; index++) {
                    const element = props.practiceList.practices[index];
                    Radar.ui.marker({
                        color: 'red',
                        scale: .5,
                        popup: {
                            text: element.name
                        },
                    }).setLngLat([element.longitude, element.lattitude]).addTo(map)
                }
                map.fitToMarkers()
            }
 

            if (props.mapConfiguration != undefined && props.mapConfiguration.RadiusFeature) {
                map.clearFeatures()
                Radar.ui.marker(
                    {
                        color: "blue",
                        scale: .75,
                    }
                ).setLngLat(props.mapConfiguration.RadiusOrigin).addTo(map)

                console.log(ZoneCoords(props.mapConfiguration.RadiusOrigin, props.mapConfiguration.Radius, 100))
                map.addPolygon({
                    type: "Feature",
                    id: 1,
                    properties: {
                        name: "radius"
                    },
                    geometry: {
                        type: "Polygon",
                        coordinates:ZoneCoords(props.mapConfiguration.RadiusOrigin, props.mapConfiguration.Radius, 50)
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

        }
      }, [props.practiceList, props.mapConfiguration, map]);

      return(
        <div id="map" style={{width: "1152px", height:600}} className="rounded-b-2xl"/>
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