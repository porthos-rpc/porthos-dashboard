import React from 'react';
import {AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip} from 'Recharts';
import './styles.css'

class Card extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        const data = [
              {name: 'Page A', value: 4000},
              {name: 'Page B', value: 3000},
              {name: 'Page C', value: 2000},
              {name: 'Page D', value: 2780},
              {name: 'Page E', value: 1890},
              {name: 'Page F', value: 2390},
              {name: 'Page G', value: 1490},
        ];

        return (
            <div className="card">
                <h3>{this.props.data.serviceName}</h3> 
                <h2>{this.props.data.methodName}</h2>
                <div className="chartAvgs">
                    <div className="chart">
                        <AreaChart width={280} height={80} data={data} margin={{top: 0, right: 0, left: 0, bottom: 0}}>
                            <Tooltip/>
                            <Area type='monotone' dataKey='value' stroke='#2c4c30' fill='#c0d0b6' />
                        </AreaChart> 
                    </div>
                    <div className="avgs">
                        <div className="avgResponseTime">{this.props.data.avgResponseTime} ms</div> 
                        <div className="avgThroughput">{this.props.data.avgThroughput} RPM</div> 
                    </div>
                </div>
                <div className="cardCommonMetric">
                    <div className="left">ResponseTime</div> <div className="right"><b>Min: 10 ms Max: 100 ms</b></div>
                </div>
                <div className="cardCommonMetric">
                    <div className="left">Throughput</div> <div className="right"><b>Min: 10 RPM Max: 100 RPM</b></div>
                </div>
            </div>
        );
    }

}

export default Card;

