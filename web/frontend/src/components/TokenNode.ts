/* eslint-disable  @typescript-eslint/no-explicit-any */
import { Node } from "@baklavajs/core";
import {INodeState} from "@baklavajs/core/dist/baklavajs-core/types";

export default class TokenNode extends Node {

    public type = "TokenNode";
    public name = this.type;

    public chaincodeName;
    public methodName;

    private counter = 0;

    public constructor(chaincodeName: string, methodName: string) {
        super();
        this.chaincodeName = chaincodeName
        this.methodName = methodName
    }

    public load(state: INodeState) {
        state.interfaces.forEach(([name, intfState]) => {
            const intf = this.addInputInterface(name);
            /* eslint-disable  @typescript-eslint/no-non-null-assertion */
            intf!.id = intfState.id;
        });
        this.counter = state.interfaces.length;
        this.chaincodeName = state.chaincodeName;
        this.methodName = state.methodName;
        super.load(state);
    }

    public save(): INodeState {

        const interfacesArr = Array.from(this.interfaces.entries()).map(([k, i]) => [k, i.save()['id'], i.save['value'] ]) as any

        const interfaces = new Array<any>();

        interfacesArr.forEach(function(iface) {
            interfaces.push({name: iface[0], id: iface[1], value: iface[2]})
        });

        const optionsArr = Array.from(this.options.entries()).map(([k, o]) => [k, o.value]) as any

        const options = new Array<any>();

        optionsArr.forEach(function(option) {
            options.push({name: option[0], value: option[1]})
        });

        const state: INodeState = {
            type: this.type,
            id: this.id,
            name: this.name,
            chaincodeName: this.chaincodeName,
            tokenId: "",
            methodName: this.methodName,
            options: options,
            state: this.state,
            interfaces: interfaces
        };
        return this.hooks.save.execute(state);
    }

    public action(action: string) {
        if (action === "Add Input") {
            this.addInputInterface("Input " + (++this.counter));
        } else if (action === "Remove Input" && this.counter > 0) {
            this.removeInterface("Input " + (this.counter--));
        }
    }

}