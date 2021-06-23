import pyautogui
import cv2
import numpy as np
import time


original  = cv2.imread("original3.png",cv2.IMREAD_GRAYSCALE)
#myScreenshot = pyautogui.screenshot(region=(610,300, 650, 590))
#original = np.array(myScreenshot)
#original = np.array(original * 255, dtype = np.uint8)

template = cv2.imread("template.png",cv2.IMREAD_GRAYSCALE)
ball = cv2.imread("ball.png",cv2.IMREAD_GRAYSCALE)

height,width = original.shape

height_2 = int(height/2)
width_2 = int(width/2)

width_rec = 300
height_rec = 300

alpha = 15

circle = np.zeros((height,width))
cv2.circle(circle,(int(width/2),int(height/2)),270,1,thickness=100)

def getROI(loc):

    rectangle = np.zeros((height,width))
    #loc=(loc[1],loc[0])
    #print("coordenates: height:" + str(loc[0]) + "width: " + str(loc[1]))

    if(loc[1] > height_2):
        if(loc[0]> width_2):
            #4 cuadrante
            #return cv2.rectangle(rectangle, (loc[0] - width_rec, loc[1] - height_rec), (loc[0],loc[1]), -1, 2)
            return original[loc[1] - height_rec:loc[1]-alpha, loc[0] - width_rec:loc[0] - alpha]
        else:
            #3 cuadrante
            #return cv2.rectangle(rectangle, (loc[0], loc[1] - height_rec), (loc[0] + width_rec,loc[1]), -1, 2)
            return original[loc[1] - height_rec : loc[1]-alpha, loc[0] + alpha:loc[0] + width_rec]
    
    if(loc[0] < width_2):
        #2 cuadrante
        #return cv2.rectangle(rectangle, loc, (loc[0] + width_rec, loc[1] + height_rec), -1, 2)
        return original[loc[1] + alpha: loc[1] + height_rec , loc[0] + alpha:loc[0] + width_rec]
    else:
        #1 cuadrante
        #return cv2.rectangle(rectangle, (loc[0]  - width_rec, loc[1]), (loc[0] + width_rec, loc[1] + height_rec), -1, 2)
        return original[loc[1] + alpha:loc[1] + height_rec, loc[0]  - width_rec:loc[0] - alpha]

input("start")
time.sleep(2)
x=0
while x < 40:
    x = x + 1
    original = np.array(pyautogui.screenshot(region=(610,300, 650, 590)))
    original = cv2.cvtColor(original, cv2.COLOR_BGR2GRAY)

    prepared_img = (circle / 255.0) * original
    prepared_img = prepared_img.astype(np.uint8)
    
    '''
    myScreenshot = pyautogui.screenshot(region=(610,300, 650, 590))
    shot = np.array(myScreenshot)
    cv2.imshow("screen",shot)
    cv2.waitKey()
    '''

    res = cv2.matchTemplate(prepared_img,template,cv2.TM_CCOEFF_NORMED)
    _, max_val, _, max_loc = cv2.minMaxLoc(res)

    '''
    cv2.circle(original,tuple((max_loc)), 13,(255,255,255), 6)
    cv2.imshow("",original)
    cv2.waitKey()
    '''
    roi = getROI(max_loc)
    '''
    prepared_roi = (roi / 255.0) * original
    prepared_roi = prepared_img.astype(np.uint8)

    cv2.imshow("",roi)
    cv2.waitKey()
    '''
    cv2.imshow("",roi)
    cv2.waitKey()
    res = cv2.matchTemplate(roi,ball,cv2.TM_CCOEFF_NORMED)
    _, max_val_ball, _, max_loc_ball = cv2.minMaxLoc(res)
    '''
    cv2.circle(roi,tuple((max_loc_ball)), 13,(255,255,255), 6)
    cv2.imshow("",roi)
    cv2.waitKey()
    '''
    
    if(max_val_ball > 0.65):
        print(max_val_ball)
        print("click")
        pyautogui.click(x=960, y=540)
    else:
        print(max_val_ball)
        print("no_click")
    time.sleep(0.15)