# learnGo

```mermaid
flowchart TD
    %% Minimal Theme System (Carbon Mono)
    classDef phase fill:#1e1e2e,stroke:#b4befe,stroke-width:1px,color:#b4befe;
    classDef nodeStyle fill:#11111b,stroke:#313244,color:#cdd6f4,stroke-width:1px;
    classDef lineStyle stroke:#45475a,stroke-width:1px;

    %% Vertical Backbone Timeline
    P1 === P2 === P3 === P4 === P5

    %% Phase 1
    P1[01 . SYNTAX FOUNDATIONS] --> 1[1_hello_world] --> 2[2_simple_value] --> 3[3_variables] --> 4[4_constants] --> 5[5_for] --> 6[6_if_else] --> 7[7_switch]

    %% Phase 2
    P2[02 . DATA STRUCTURES] --> 8[8_arrays] --> 9[9_slices] --> 10[10_map] --> 11[11_range] --> 15[15_pointers]

    %% Phase 3
    P3[03 . FUNCTIONAL BLOCKS] --> 12[12_function] --> 13[13_variadic] --> 14[14_closures]

    %% Phase 4
    P4[04 . GO TYPE SYSTEM] --> 16[16_structs] --> 18[18_enums] --> 17[17_interface] --> 19[19_generics]

    %% Phase 5
    P5[05 . CONCURRENCY & I/O] --> 20[20_goroutines] --> 21[21_channels] --> 22[22_mutex] --> 23[23_files] --> 24[24_package]

    %% Class Bindings
    class P1,P2,P3,P4,P5 phase;
    class 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24 nodeStyle;